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

type portcontrollerhandler struct {
	kubeService services.IKubeService
}

func (h *portcontrollerhandler) router() chi.Router {
	r := chi.NewRouter()

	r.Route("/service", func(r chi.Router) {
		r.Get("/*", h.defaultHandler)
	})
	r.Route("/deployments", func(r chi.Router) {
		r.Post("/list", h.deploymentsHandler)
		r.Post("/update", h.deploymentUpdateHandler)
	})
	r.Route("/forward", func(r chi.Router) {
		r.Post("/", h.portforwardHandler)
	})
	r.Route("/namespaces", func(r chi.Router) {
		r.Get("/*", h.namespaceHandler)
	})
	r.Route("/nodes", func(r chi.Router) {
		r.Get("/*", h.kubeNodesHandler)
	})
	r.Route("/pod", func(r chi.Router) {
		r.Post("/restart", h.kubePodHandler)
	})

	return r
}

func (h *portcontrollerhandler) kubePodHandler(w http.ResponseWriter, r *http.Request) {
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

func (h *portcontrollerhandler) defaultHandler(w http.ResponseWriter, r *http.Request) {

	var pathInfo = strings.Split(r.RequestURI, "/")
	if pathInfo != nil && len(pathInfo) > 2 {
		resp := h.kubeService.GetByNamespace(pathInfo[len(pathInfo)-1])
		render.JSON(w, r, resp)
	}
}

func (h *portcontrollerhandler) namespaceHandler(w http.ResponseWriter, r *http.Request) {

	var pathInfo = strings.Split(r.RequestURI, "/")
	if pathInfo != nil && len(pathInfo) > 2 {
		resp := h.kubeService.GetNamespaces()
		render.JSON(w, r, resp)
	}
}

func (h *portcontrollerhandler) deploymentsHandler(w http.ResponseWriter, r *http.Request) {
	reqBoyd, _ := ioutil.ReadAll(r.Body)
	requestObj := struct {
		Namespace string `json:"namespace"`
	}{}
	_ = json.Unmarshal(reqBoyd, &requestObj)

	resp := h.kubeService.GetDeployments(requestObj.Namespace)

	render.JSON(w, r, resp)
}

func (h *portcontrollerhandler) deploymentUpdateHandler(w http.ResponseWriter, r *http.Request) {
	reqBoyd, _ := ioutil.ReadAll(r.Body)
	requestObj := struct {
		Namespace      string `json:"namespace"`
		ContainerName  string `json:"containername"`
		NewImage       string `json:"newimage"`
		DeploymentName string `json:"deploymentname"`
	}{}
	_ = json.Unmarshal(reqBoyd, &requestObj)

	resp := h.kubeService.UpdateImageForDeployment(requestObj.DeploymentName, requestObj.ContainerName, requestObj.NewImage, requestObj.Namespace)

	respObject := struct {
		Success bool `json:"success"`
	}{
		Success: resp,
	}
	render.JSON(w, r, respObject)
}

func (h *portcontrollerhandler) portforwardHandler(w http.ResponseWriter, r *http.Request) {
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

func (h *portcontrollerhandler) kubeNodesHandler(w http.ResponseWriter, r *http.Request) {
	resp := h.kubeService.GetNodes()
	render.JSON(w, r, resp)
}
