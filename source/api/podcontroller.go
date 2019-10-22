package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"../services"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

type podcontrollerhandler struct {
	kubeService services.IKubeService
}

func (h *podcontrollerhandler) router() chi.Router {
	r := chi.NewRouter()

	r.Route("/info", func(r chi.Router) {
		r.Get("/*", h.getByNamespaceHandler)
	})

	r.Route("/forward", func(r chi.Router) {
		r.Post("/", h.portforwardHandler)

	})

	r.Route("/restart", func(r chi.Router) {
		r.Post("/", h.kubePodHandler)

	})

	return r
}

func (h *podcontrollerhandler) kubePodHandler(w http.ResponseWriter, r *http.Request) {
	reqBoyd, _ := ioutil.ReadAll(r.Body)
	requestObj := struct {
		Namespace string `json:"namespace"`
		PodName   string `json:"podname"`
	}{}
	_ = json.Unmarshal(reqBoyd, &requestObj)

	resp := h.kubeService.RestartPod(requestObj.PodName, requestObj.Namespace)

	respObject := struct {
		Success bool `json:"success"`
	}{
		Success: resp,
	}

	render.JSON(w, r, respObject)
}

func (h *podcontrollerhandler) getByNamespaceHandler(w http.ResponseWriter, r *http.Request) {

	var pathInfo = strings.Split(r.RequestURI, "/")
	if pathInfo != nil && len(pathInfo) > 2 {
		resp := h.kubeService.GetByNamespace(pathInfo[len(pathInfo)-1])
		render.JSON(w, r, resp)
	}
}

func (h *podcontrollerhandler) portforwardHandler(w http.ResponseWriter, r *http.Request) {
	reqBoyd, _ := ioutil.ReadAll(r.Body)
	requestObj := struct {
		LocalPort       string `json:"localport"`
		DestinationPort string `json:"destinationport"`
		PodName         string `json:"podname"`
		Namespace       string `json:"namespace"`
	}{}
	_ = json.Unmarshal(reqBoyd, &requestObj)

	resp := h.kubeService.GetPortForwardCommand(requestObj.PodName, requestObj.Namespace,
		requestObj.DestinationPort, requestObj.LocalPort)

	render.JSON(w, r, resp)
}
