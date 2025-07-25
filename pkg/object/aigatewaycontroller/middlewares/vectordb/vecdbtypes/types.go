/*
 * Copyright (c) 2017, The Easegress Authors
 * All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package vecdbtypes

import (
	"context"
	"errors"
)

var ErrSimilaritySearchNotFound = errors.New("not found a result that matches the query in vector database")

type (
	// VectorDB is the interface for vector database middleware.
	VectorDB interface {
		CreateSchema(ctx context.Context, options ...Option) (VectorHandler, error)
	}

	VectorHandler interface {
		SimilaritySearch(ctx context.Context, options ...HandlerSearchOption) ([]map[string]any, error)
		InsertDocuments(ctx context.Context, doc []map[string]any, options ...HandlerInsertOption) ([]string, error)
	}

	// CommonSpec defines the specification for a vector database middleware.
	CommonSpec struct {
		Type           string  `json:"type"`
		Threshold      float64 `json:"threshold" jsonschema:"required"`
		CollectionName string  `json:"collectionName" jsonschema:"required"`
	}
)
