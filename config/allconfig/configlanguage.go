// Copyright 2023 The Hugo Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package allconfig

import (
	"time"

	"github.com/gohugoio/hugo/common/urls"
	"github.com/gohugoio/hugo/config"
	"github.com/gohugoio/hugo/langs"
)

type ConfigLanguage struct {
	config     *Config
	baseConfig config.BaseConfig

	m        *Configs
	language *langs.Language
}

func (c ConfigLanguage) Language() *langs.Language {
	return c.language
}

func (c ConfigLanguage) Languages() langs.Languages {
	return c.m.Languages
}

func (c ConfigLanguage) LanguagesDefaultFirst() langs.Languages {
	return c.m.LanguagesDefaultFirst
}

func (c ConfigLanguage) BaseURL() urls.BaseURL {
	return c.config.C.BaseURL
}

func (c ConfigLanguage) BaseURLLiveReload() urls.BaseURL {
	return c.config.C.BaseURLLiveReload
}

func (c ConfigLanguage) Environment() string {
	return c.config.Environment
}

func (c ConfigLanguage) IsMultihost() bool {
	return c.m.IsMultihost
}

func (c ConfigLanguage) IsMultiLingual() bool {
	return len(c.m.Languages) > 1
}

func (c ConfigLanguage) TemplateMetrics() bool {
	return c.config.TemplateMetrics
}

func (c ConfigLanguage) TemplateMetricsHints() bool {
	return c.config.TemplateMetricsHints
}

func (c ConfigLanguage) IsLangDisabled(lang string) bool {
	return c.config.C.DisabledLanguages[lang]
}

func (c ConfigLanguage) IgnoredErrors() map[string]bool {
	return c.config.C.IgnoredErrors
}

func (c ConfigLanguage) NoBuildLock() bool {
	return c.config.NoBuildLock
}

func (c ConfigLanguage) NewContentEditor() string {
	return c.config.NewContentEditor
}

func (c ConfigLanguage) Timeout() time.Duration {
	return c.config.C.Timeout
}

func (c ConfigLanguage) BaseConfig() config.BaseConfig {
	return c.baseConfig
}

func (c ConfigLanguage) Dirs() config.CommonDirs {
	return c.config.CommonDirs
}

func (c ConfigLanguage) DirsBase() config.CommonDirs {
	return c.m.Base.CommonDirs
}

func (c ConfigLanguage) Quiet() bool {
	return c.m.Base.Internal.Quiet
}

// GetConfigSection is mostly used in tests. The switch statement isn't complete, but what's in use.
func (c ConfigLanguage) GetConfigSection(s string) any {
	switch s {
	case "security":
		return c.config.Security
	case "build":
		return c.config.Build
	case "frontmatter":
		return c.config.Frontmatter
	case "caches":
		return c.config.Caches
	case "markup":
		return c.config.Markup
	case "mediaTypes":
		return c.config.MediaTypes.Config
	case "outputFormats":
		return c.config.OutputFormats.Config
	case "permalinks":
		return c.config.Permalinks
	case "minify":
		return c.config.Minify
	case "activeModules":
		return c.m.Modules
	case "deployment":
		return c.config.Deployment
	default:
		panic("not implemented: " + s)
	}
}

func (c ConfigLanguage) GetConfig() any {
	return c.config
}

func (c ConfigLanguage) CanonifyURLs() bool {
	return c.config.CanonifyURLs
}

func (c ConfigLanguage) IsUglyURLs(section string) bool {
	return c.config.C.IsUglyURLSection(section)
}

func (c ConfigLanguage) IgnoreFile(s string) bool {
	return c.config.C.IgnoreFile(s)
}

func (c ConfigLanguage) DisablePathToLower() bool {
	return c.config.DisablePathToLower
}

func (c ConfigLanguage) RemovePathAccents() bool {
	return c.config.RemovePathAccents
}

func (c ConfigLanguage) DefaultContentLanguage() string {
	return c.config.DefaultContentLanguage
}

func (c ConfigLanguage) DefaultContentLanguageInSubdir() bool {
	return c.config.DefaultContentLanguageInSubdir
}

func (c ConfigLanguage) SummaryLength() int {
	return c.config.SummaryLength
}

func (c ConfigLanguage) BuildExpired() bool {
	return c.config.BuildExpired
}

func (c ConfigLanguage) BuildFuture() bool {
	return c.config.BuildFuture
}

func (c ConfigLanguage) BuildDrafts() bool {
	return c.config.BuildDrafts
}

func (c ConfigLanguage) Running() bool {
	return c.config.Internal.Running
}

func (c ConfigLanguage) PrintUnusedTemplates() bool {
	return c.config.PrintUnusedTemplates
}

func (c ConfigLanguage) EnableMissingTranslationPlaceholders() bool {
	return c.config.EnableMissingTranslationPlaceholders
}

func (c ConfigLanguage) LogI18nWarnings() bool {
	return c.config.LogI18nWarnings
}

func (c ConfigLanguage) CreateTitle(s string) string {
	return c.config.C.CreateTitle(s)
}

func (c ConfigLanguage) Paginate() int {
	return c.config.Paginate
}

func (c ConfigLanguage) PaginatePath() string {
	return c.config.PaginatePath
}

func (c ConfigLanguage) StaticDirs() []string {
	return c.config.staticDirs()
}
