package jobs

import (
    "errors"
    "sync"
    "time"
)

type JobStatus struct {
    Status string `json:"status"`
    JobID  string `json:"job_id"`
    Error  []struct {
        StoreID string `json:"store_id"`
        Error   string `json:"error"`
    } `json:"error,omitempty"`
}

var (
    jobCounter   int
    jobs         = make(map[string]*JobStatus)
    jobsMutex    sync.Mutex
)

func SubmitJob(visits []Visit) (string, error) {
    jobsMutex.Lock()
    defer jobsMutex.Unlock()
    
    jobID := fmt.Sprintf("job-%d", jobCounter)
    jobCounter++

    jobs[jobID] = &JobStatus{Status: "ongoing", JobID: jobID}
    
    go processJob(jobID, visits)
    
    return jobID, nil
}

func GetJobStatus(jobID string) (*JobStatus, error) {
    jobsMutex.Lock()
    defer jobsMutex.Unlock()
    
    if job, exists := jobs[jobID]; exists {
        return job, nil
    }
    return nil, errors.New("job not found")
}

func processJob(jobID string, visits []Visit) {
    for _, visit := range visits {
        if err := ProcessImages(visit.StoreID, visit.ImageURLs); err != nil {
            jobs[jobID].Status = "failed"
            jobs[jobID].Error = append(jobs[jobID].Error, struct {
                StoreID string `json:"store_id"`
                Error   string `json:"error"`
            }{StoreID: visit.StoreID, Error: err.Error()})
            return
        }
    }
    jobs[jobID].Status = "completed"
}
