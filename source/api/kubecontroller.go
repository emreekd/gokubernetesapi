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

type kubecontrollerhandler struct {
	kubeService services.IKubeService
}

func (h *kubecontrollerhandler) router() chi.Router {
	r := chi.NewRouter()

	r.Route("/deployments", func(r chi.Router) {
		r.Post("/list", h.deploymentsHandler)
		r.Post("/update", h.deploymentUpdateHandler)
	})
	r.Route("/namespaces", func(r chi.Router) {
		r.Get("/*", h.namespaceHandler)
	})
	r.Route("/nodes", func(r chi.Router) {
		r.Get("/*", h.kubeNodesHandler)
	})

	return r
}

func (h *kubecontrollerhandler) namespaceHandler(w http.ResponseWriter, r *http.Request) {

	var pathInfo = strings.Split(r.RequestURI, "/")
	if pathInfo != nil && len(pathInfo) > 2 {
		resp := h.kubeService.GetNamespaces()
		render.JSON(w, r, resp)
	}
}

func (h *kubecontrollerhandler) deploymentsHandler(w http.ResponseWriter, r *http.Request) {
	reqBoyd, _ := ioutil.ReadAll(r.Body)
	requestObj := struct {
		Namespace string `json:"namespace"`
	}{}
	_ = json.Unmarshal(reqBoyd, &requestObj)

	resp := h.kubeService.GetDeployments(requestObj.Namespace)

	render.JSON(w, r, resp)
}

func (h *kubecontrollerhandler) deploymentUpdateHandler(w http.ResponseWriter, r *http.Request) {
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

func (h *kubecontrollerhandler) kubeNodesHandler(w http.ResponseWriter, r *http.Request) {
	resp := h.kubeService.GetNodes()
	render.JSON(w, r, resp)
}
