package downloader

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

func isValidDir(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	if !info.IsDir() {
		return false
	}
	return true
}

func FromUrlGetFileName(url string) string {
	length := len(url)
	lastPos := strings.LastIndex(url, `/`)
	for lastPos == length-1 {
		lastPos = strings.LastIndex(url[:length-1], `/`)
		length--
	}
	return url[lastPos+1 : length]
}

type urlInfo struct {
	name string
	size uint64
	url  string
}

type httpStatusPair struct {
	statusCode int
	status     string
}

type worker struct {
	downloadDir    string
	targetInfo     urlInfo
	errorChan      *chan error
	finishChan     *chan bool
	httpStatusChan *chan httpStatusPair
	receiveChan    *chan uint64
	totalChan      *chan uint64
}

func newWorker(downloadDir string, url string) *worker {
	errorChan := make(chan error)
	finishChan := make(chan bool)
	httpStatusChan := make(chan httpStatusPair)
	receiveChan := make(chan uint64)
	totalChan := make(chan uint64)

	return &worker{
		downloadDir:    downloadDir,
		errorChan:      &errorChan,
		finishChan:     &finishChan,
		httpStatusChan: &httpStatusChan,
		receiveChan:    &receiveChan,
		totalChan:      &totalChan,
		targetInfo: urlInfo{
			name: FromUrlGetFileName(url),
			url:  url,
		},
	}
}

func (w *worker) Write(p []byte) (int, error) {
	l := len(p)
	*w.receiveChan <- uint64(l)
	return l, nil
}

func (w *worker) closeAllChannel() {
	close(*w.errorChan)
	close(*w.finishChan)
	close(*w.httpStatusChan)
	close(*w.receiveChan)
	close(*w.totalChan)
}

func (w *worker) getRequest() {
	defer w.closeAllChannel()
	defer func(ch chan bool) { ch <- true }(*w.finishChan)
	fileName := fmt.Sprintf("%s/%s", w.downloadDir, w.targetInfo.name)
	tmpFileName := fileName + ".part"
	f, err := os.Create(tmpFileName)
	defer func(fileName string) {
		if _, err := os.Stat(tmpFileName); err == nil {
			_ = os.Remove(tmpFileName)
		}
	}(tmpFileName)
	if err != nil {
		*w.errorChan <- err
		return
	}
	client := &http.Client{}
	req, err := http.NewRequest("GET", w.targetInfo.url, nil)
	if err != nil {
		*w.errorChan <- err
		return
	}
	res, err := client.Do(req)
	defer func() { _ = res.Body.Close() }()
	if err != nil {
		*w.errorChan <- err
		return
	}

	s, _ := strconv.Atoi(strings.Join(res.Header[`Content-Length`], ``))
	*w.totalChan <- uint64(s)
	*w.httpStatusChan <- httpStatusPair{status: res.Status, statusCode: res.StatusCode}
	if res.StatusCode > 200 {
		*w.finishChan <- true
		return
	}
	_, err = io.Copy(f, io.TeeReader(res.Body, w))
	if err != nil {
		*w.errorChan <- err
		return
	}
	err = os.Rename(tmpFileName, fileName)
	if err != nil {
		*w.errorChan <- err
		return
	}
}

type watcher struct {
	name string
	url  string
	wo   *worker

	wSize          uint64
	wTotalSize     uint64
	workError      error
	isFinished     bool
	workHttpStatus httpStatusPair

	errorChan      *chan error
	finishChan     *chan bool
	httpStatusChan *chan httpStatusPair
	receiveChan    *chan uint64
	totalChan      *chan uint64
	closedWG       *sync.WaitGroup
}

func newWatcher(downloadDir string, url string, wg *sync.WaitGroup) *watcher {
	wo := newWorker(downloadDir, url)
	return &watcher{
		name:           wo.targetInfo.name,
		url:            wo.targetInfo.url,
		wo:             wo,
		errorChan:      wo.errorChan,
		finishChan:     wo.finishChan,
		httpStatusChan: wo.httpStatusChan,
		receiveChan:    wo.receiveChan,
		totalChan:      wo.totalChan,
		closedWG:       wg,
	}
}

func (w *watcher) updateWorkInfo() {
	defer w.closedWG.Done()
	for {
		select {
		case w.workError = <-*w.errorChan:
			//fmt.Println("ERROR:", w.name, w.workError)
		case w.isFinished = <-*w.finishChan:
			//fmt.Println("FINISHED:", w.name, w.isFinished)
			return
		case w.workHttpStatus = <-*w.httpStatusChan:
		case size := <-*w.receiveChan:
			w.wSize += size
		case w.wTotalSize = <-*w.totalChan:
		}
		//time.Sleep(time.Millisecond * 100)
	}
}

type Downloader struct {
	downloadDir   string
	WatchDuration time.Duration
	watcherWG     sync.WaitGroup
	allExitChan   chan bool
	allExitWG     sync.WaitGroup

	watchers []*watcher
	spaces   string
}

func (d *Downloader) watchOnceTime() {
	var toPrint []string
	toPrint = append(toPrint, fmt.Sprintf("name%-"+d.spaces+"sstatus\n", " "))
	for _, wa := range d.watchers {
		var reqLine string
		reqLine = fmt.Sprintf("%-"+d.spaces+"s", wa.name)

		// Check if error occurred.
		//fmt.Println("CHECK wa.workError", wa.workError)
		if wa.workError != nil {
			reqLine += fmt.Sprintf("%s\n", wa.workError)
		} else if wa.workHttpStatus.statusCode > 200 {
			reqLine += fmt.Sprintf(" <%s>\n", wa.workHttpStatus.status)
		} else if wa.wTotalSize == 0 && !wa.isFinished {
			reqLine += fmt.Sprintf(" %4s %s\n", "-", wa.workHttpStatus.status)
		} else {
			reqLine += fmt.Sprintf(" %3d%%[%d]\n", 100*wa.wSize/wa.wTotalSize, wa.wSize)
		}
		toPrint = append(toPrint, reqLine)
	}

	// Return to start line.
	lines := len(toPrint)
	for i := 0; i < lines; i++ {
		toPrint = append(toPrint, fmt.Sprintf("\033[1A"))
	}

	//Print.
	for _, s := range toPrint {
		fmt.Printf("%s", s)
	}
}

func (d *Downloader) watchAllWatcher() {
	d.allExitWG.Add(1)
	for {
		d.watchOnceTime()
		select {
		case <-d.allExitChan:
			d.watchOnceTime()
			// Restore terminal line position and exit.
			for i := 0; i <= len(d.watchers); i++ {
				fmt.Printf("\n")
			}
			d.allExitWG.Done()
			return
		case <-time.After(d.WatchDuration):
			continue
		}
	}

}

func (d *Downloader) startWorks() {
	for _, wa := range d.watchers {
		go wa.wo.getRequest()
	}
}

func (d *Downloader) startWatchers() {
	for _, re := range d.watchers {
		go re.updateWorkInfo()
	}
	go d.watchAllWatcher()
}

// Start starts the downloader d to download.
func (d *Downloader) Start() {
	d.startWatchers()
	d.startWorks()
	d.watcherWG.Wait()
	d.allExitChan <- true
	close(d.allExitChan)
	d.allExitWG.Wait()
}

// NewDownloader returns a downloader pointer with given option args:
// Download files and save in dir downloadDir.
// Download files from urls in string slice urls.
func NewDownloader(downloadDir string, urls []string) (*Downloader, error) {
	if !isValidDir(downloadDir) {
		return nil, errors.New("Not a valid directory" + downloadDir)
	}
	r := Downloader{
		downloadDir:   downloadDir,
		WatchDuration: time.Millisecond * 300,
		allExitChan:   make(chan bool),
	}
	nameLengthMax := 0
	for _, url := range urls {
		r.watcherWG.Add(1)
		wa := newWatcher(downloadDir, url, &r.watcherWG)
		r.watchers = append(r.watchers, wa)
		if nameLengthMax < len(wa.name) {
			nameLengthMax = len(wa.name)
		}
	}
	r.spaces = strconv.Itoa(nameLengthMax + 2)
	return &r, nil
}
