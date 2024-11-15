package handlers

import (
    "encoding/json"
    "net/http"
    "retail-pulse-backend/jobs"
    "retail-pulse-backend/store"
    "retail-pulse-backend/utils"
)

type SubmitJobRequest struct {
    Count  int      `json:"count"`
    Visits []Visit  `json:"visits"`
}

type Visit struct {
    StoreID   string   `json:"store_id"`
    ImageURLs []string `json:"image_url"`
    VisitTime string   `json:"visit_time"`
}

func SubmitJob(w http.ResponseWriter, r *http.Request) {
    var req SubmitJobRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.Count != len(req.Visits) {
        utils.JSONError(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    // Validate each store_id
    for _, visit := range req.Visits {
        if !store.IsValidStoreID(visit.StoreID) {
            utils.JSONError(w, "Invalid store ID: "+visit.StoreID, http.StatusBadRequest)
            return
        }
    }

    // Submit job
    jobID, err := jobs.SubmitJob(req.Visits)
    if err != nil {
        utils.JSONError(w, "Failed to submit job", http.StatusInternalServerError)
        return
    }

    utils.JSONResponse(w, map[string]interface{}{"job_id": jobID}, http.StatusCreated)
}

func GetJobStatus(w http.ResponseWriter, r *http.Request) {
    jobID := r.URL.Query().Get("jobid")
    status, err := jobs.GetJobStatus(jobID)
    if err != nil {
        utils.JSONError(w, "Job ID not found", http.StatusBadRequest)
        return
    }
    utils.JSONResponse(w, status, http.StatusOK)
}
