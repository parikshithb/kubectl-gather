// SPDX-FileCopyrightText: The kubectl-gather authors
// SPDX-License-Identifier: Apache-2.0

package gather

import (
	"net/http"

	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/rest"
)

type gatherBackend struct {
	g  *Gatherer
	wq *WorkQueue
}

func (b *gatherBackend) Config() *rest.Config {
	return b.g.config
}

func (b *gatherBackend) HTTPClient() *http.Client {
	return b.g.httpClient
}

func (b *gatherBackend) Options() *Options {
	return b.g.opts
}

func (b *gatherBackend) Output() *OutputDirectory {
	return &b.g.output
}

func (b *gatherBackend) Queue(work WorkFunc) {
	b.wq.Queue(work)
}

func (b *gatherBackend) GatherResource(gvr schema.GroupVersionResource, name types.NamespacedName) {
	b.g.gatherResource(gvr, name)
}
