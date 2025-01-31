package assets

import (
	"context"
	"strings"

	"github.com/zitadel/zitadel/internal/api/authz"
	"github.com/zitadel/zitadel/internal/command"
	"github.com/zitadel/zitadel/internal/domain"
	"github.com/zitadel/zitadel/internal/id"
	"github.com/zitadel/zitadel/internal/query"
)

func (h *Handler) UploadDefaultLabelPolicyLogo() Uploader {
	return &labelPolicyLogoUploader{h.idGenerator, false, true, []string{"image/"}, 1 << 19}
}

func (h *Handler) UploadDefaultLabelPolicyLogoDark() Uploader {
	return &labelPolicyLogoUploader{h.idGenerator, true, true, []string{"image/"}, 1 << 19}
}

func (h *Handler) UploadOrgLabelPolicyLogo() Uploader {
	return &labelPolicyLogoUploader{h.idGenerator, false, false, []string{"image/"}, 1 << 19}
}

func (h *Handler) UploadOrgLabelPolicyLogoDark() Uploader {
	return &labelPolicyLogoUploader{h.idGenerator, true, false, []string{"image/"}, 1 << 19}
}

type labelPolicyLogoUploader struct {
	idGenerator   id.Generator
	darkMode      bool
	defaultPolicy bool
	contentTypes  []string
	maxSize       int64
}

func (l *labelPolicyLogoUploader) ContentTypeAllowed(contentType string) bool {
	for _, ct := range l.contentTypes {
		if strings.HasPrefix(contentType, ct) {
			return true
		}
	}
	return false
}

func (l *labelPolicyLogoUploader) MaxFileSize() int64 {
	return l.maxSize
}

func (l *labelPolicyLogoUploader) ObjectName(_ authz.CtxData) (string, error) {
	suffixID, err := l.idGenerator.Next()
	if err != nil {
		return "", err
	}
	prefix := domain.LabelPolicyLogoPath
	if l.darkMode {
		return prefix + "-" + domain.Dark + "-" + suffixID, nil
	}
	return prefix + "-" + suffixID, nil
}

func (l *labelPolicyLogoUploader) BucketName(ctxData authz.CtxData) string {
	if l.defaultPolicy {
		return domain.IAMID
	}
	return ctxData.OrgID
}

func (l *labelPolicyLogoUploader) Callback(ctx context.Context, info *domain.AssetInfo, orgID string, commands *command.Commands) error {
	if l.defaultPolicy {
		if l.darkMode {
			_, err := commands.AddLogoDarkDefaultLabelPolicy(ctx, info.Key)
			return err
		}
		_, err := commands.AddLogoDefaultLabelPolicy(ctx, info.Key)
		return err
	}
	if l.darkMode {
		_, err := commands.AddLogoDarkLabelPolicy(ctx, orgID, info.Key)
		return err
	}
	_, err := commands.AddLogoLabelPolicy(ctx, orgID, info.Key)
	return err
}

func (h *Handler) GetDefaultLabelPolicyLogo() Downloader {
	return &labelPolicyLogoDownloader{query: h.query, darkMode: false, defaultPolicy: true, preview: false}
}

func (h *Handler) GetDefaultLabelPolicyLogoDark() Downloader {
	return &labelPolicyLogoDownloader{query: h.query, darkMode: true, defaultPolicy: true, preview: false}
}

func (h *Handler) GetPreviewDefaultLabelPolicyLogo() Downloader {
	return &labelPolicyLogoDownloader{query: h.query, darkMode: false, defaultPolicy: true, preview: true}
}

func (h *Handler) GetPreviewDefaultLabelPolicyLogoDark() Downloader {
	return &labelPolicyLogoDownloader{query: h.query, darkMode: true, defaultPolicy: true, preview: true}
}

func (h *Handler) GetOrgLabelPolicyLogo() Downloader {
	return &labelPolicyLogoDownloader{query: h.query, darkMode: false, defaultPolicy: false, preview: false}
}

func (h *Handler) GetOrgLabelPolicyLogoDark() Downloader {
	return &labelPolicyLogoDownloader{query: h.query, darkMode: true, defaultPolicy: false, preview: false}
}

func (h *Handler) GetPreviewOrgLabelPolicyLogo() Downloader {
	return &labelPolicyLogoDownloader{query: h.query, darkMode: false, defaultPolicy: false, preview: true}
}

func (h *Handler) GetPreviewOrgLabelPolicyLogoDark() Downloader {
	return &labelPolicyLogoDownloader{query: h.query, darkMode: true, defaultPolicy: false, preview: true}
}

type labelPolicyLogoDownloader struct {
	query         *query.Queries
	darkMode      bool
	defaultPolicy bool
	preview       bool
}

func (l *labelPolicyLogoDownloader) ObjectName(ctx context.Context, path string) (string, error) {
	policy, err := getLabelPolicy(ctx, l.defaultPolicy, l.preview, l.query)
	if err != nil {
		return "", nil
	}
	if l.darkMode {
		return policy.Dark.LogoURL, nil
	}
	return policy.Light.LogoURL, nil
}

func (l *labelPolicyLogoDownloader) BucketName(ctx context.Context, id string) string {
	return getLabelPolicyBucketName(ctx, l.defaultPolicy, l.preview, l.query)
}

func (h *Handler) UploadDefaultLabelPolicyIcon() Uploader {
	return &labelPolicyIconUploader{h.idGenerator, false, true, []string{"image/"}, 1 << 19}
}

func (h *Handler) UploadDefaultLabelPolicyIconDark() Uploader {
	return &labelPolicyIconUploader{h.idGenerator, true, true, []string{"image/"}, 1 << 19}
}

func (h *Handler) UploadOrgLabelPolicyIcon() Uploader {
	return &labelPolicyIconUploader{h.idGenerator, false, false, []string{"image/"}, 1 << 19}
}

func (h *Handler) UploadOrgLabelPolicyIconDark() Uploader {
	return &labelPolicyIconUploader{h.idGenerator, true, false, []string{"image/"}, 1 << 19}
}

type labelPolicyIconUploader struct {
	idGenerator   id.Generator
	darkMode      bool
	defaultPolicy bool
	contentTypes  []string
	maxSize       int64
}

func (l *labelPolicyIconUploader) ContentTypeAllowed(contentType string) bool {
	for _, ct := range l.contentTypes {
		if strings.HasPrefix(contentType, ct) {
			return true
		}
	}
	return false
}

func (l *labelPolicyIconUploader) MaxFileSize() int64 {
	return l.maxSize
}

func (l *labelPolicyIconUploader) ObjectName(_ authz.CtxData) (string, error) {
	suffixID, err := l.idGenerator.Next()
	if err != nil {
		return "", err
	}
	prefix := domain.LabelPolicyIconPath
	if l.darkMode {
		return prefix + "-" + domain.Dark + "-" + suffixID, nil
	}
	return prefix + "-" + suffixID, nil
}

func (l *labelPolicyIconUploader) BucketName(ctxData authz.CtxData) string {
	if l.defaultPolicy {
		return domain.IAMID
	}
	return ctxData.OrgID
}

func (l *labelPolicyIconUploader) Callback(ctx context.Context, info *domain.AssetInfo, orgID string, commands *command.Commands) error {
	if l.defaultPolicy {
		if l.darkMode {
			_, err := commands.AddIconDarkDefaultLabelPolicy(ctx, info.Key)
			return err
		}
		_, err := commands.AddIconDefaultLabelPolicy(ctx, info.Key)
		return err
	}

	if l.darkMode {
		_, err := commands.AddIconDarkLabelPolicy(ctx, orgID, info.Key)
		return err
	}
	_, err := commands.AddIconLabelPolicy(ctx, orgID, info.Key)
	return err
}

func (h *Handler) GetDefaultLabelPolicyIcon() Downloader {
	return &labelPolicyIconDownloader{query: h.query, darkMode: false, defaultPolicy: true, preview: false}
}

func (h *Handler) GetDefaultLabelPolicyIconDark() Downloader {
	return &labelPolicyIconDownloader{query: h.query, darkMode: true, defaultPolicy: true, preview: false}
}

func (h *Handler) GetPreviewDefaultLabelPolicyIcon() Downloader {
	return &labelPolicyIconDownloader{query: h.query, darkMode: false, defaultPolicy: true, preview: true}
}

func (h *Handler) GetPreviewDefaultLabelPolicyIconDark() Downloader {
	return &labelPolicyIconDownloader{query: h.query, darkMode: true, defaultPolicy: true, preview: true}
}

func (h *Handler) GetOrgLabelPolicyIcon() Downloader {
	return &labelPolicyIconDownloader{query: h.query, darkMode: false, defaultPolicy: false, preview: false}
}

func (h *Handler) GetOrgLabelPolicyIconDark() Downloader {
	return &labelPolicyIconDownloader{query: h.query, darkMode: true, defaultPolicy: false, preview: false}
}

func (h *Handler) GetPreviewOrgLabelPolicyIcon() Downloader {
	return &labelPolicyIconDownloader{query: h.query, darkMode: false, defaultPolicy: false, preview: true}
}

func (h *Handler) GetPreviewOrgLabelPolicyIconDark() Downloader {
	return &labelPolicyIconDownloader{query: h.query, darkMode: true, defaultPolicy: false, preview: true}
}

type labelPolicyIconDownloader struct {
	query         *query.Queries
	darkMode      bool
	defaultPolicy bool
	preview       bool
}

func (l *labelPolicyIconDownloader) ObjectName(ctx context.Context, path string) (string, error) {
	policy, err := getLabelPolicy(ctx, l.defaultPolicy, l.preview, l.query)
	if err != nil {
		return "", nil
	}
	if l.darkMode {
		return policy.Dark.IconURL, nil
	}
	return policy.Light.IconURL, nil
}

func (l *labelPolicyIconDownloader) BucketName(ctx context.Context, id string) string {
	return getLabelPolicyBucketName(ctx, l.defaultPolicy, l.preview, l.query)
}

func (h *Handler) UploadDefaultLabelPolicyFont() Uploader {
	return &labelPolicyFontUploader{h.idGenerator, true, []string{"font/", "application/octet-stream"}, 1 << 19}
}

func (h *Handler) UploadOrgLabelPolicyFont() Uploader {
	return &labelPolicyFontUploader{h.idGenerator, false, []string{"font/", "application/octet-stream"}, 1 << 19}
}

type labelPolicyFontUploader struct {
	idGenerator   id.Generator
	defaultPolicy bool
	contentTypes  []string
	maxSize       int64
}

func (l *labelPolicyFontUploader) ContentTypeAllowed(contentType string) bool {
	for _, ct := range l.contentTypes {
		if strings.HasPrefix(contentType, ct) {
			return true
		}
	}
	return false
}

func (l *labelPolicyFontUploader) MaxFileSize() int64 {
	return l.maxSize
}

func (l *labelPolicyFontUploader) ObjectName(_ authz.CtxData) (string, error) {
	suffixID, err := l.idGenerator.Next()
	if err != nil {
		return "", err
	}
	prefix := domain.LabelPolicyFontPath
	return prefix + "-" + suffixID, nil
}

func (l *labelPolicyFontUploader) BucketName(ctxData authz.CtxData) string {
	if l.defaultPolicy {
		return domain.IAMID
	}
	return ctxData.OrgID
}

func (l *labelPolicyFontUploader) Callback(ctx context.Context, info *domain.AssetInfo, orgID string, commands *command.Commands) error {
	if l.defaultPolicy {
		_, err := commands.AddFontDefaultLabelPolicy(ctx, info.Key)
		return err
	}
	_, err := commands.AddFontLabelPolicy(ctx, orgID, info.Key)
	return err
}

func (h *Handler) GetDefaultLabelPolicyFont() Downloader {
	return &labelPolicyFontDownloader{query: h.query, defaultPolicy: true, preview: false}
}

func (h *Handler) GetPreviewDefaultLabelPolicyFont() Downloader {
	return &labelPolicyFontDownloader{query: h.query, defaultPolicy: true, preview: true}
}

func (h *Handler) GetOrgLabelPolicyFont() Downloader {
	return &labelPolicyFontDownloader{query: h.query, defaultPolicy: false, preview: false}
}

func (h *Handler) GetPreviewOrgLabelPolicyFont() Downloader {
	return &labelPolicyFontDownloader{query: h.query, defaultPolicy: true, preview: true}
}

type labelPolicyFontDownloader struct {
	query         *query.Queries
	defaultPolicy bool
	preview       bool
}

func (l *labelPolicyFontDownloader) ObjectName(ctx context.Context, path string) (string, error) {
	policy, err := getLabelPolicy(ctx, l.defaultPolicy, l.preview, l.query)
	if err != nil {
		return "", nil
	}
	return policy.FontURL, nil
}

func (l *labelPolicyFontDownloader) BucketName(ctx context.Context, id string) string {
	return getLabelPolicyBucketName(ctx, l.defaultPolicy, l.preview, l.query)
}

func getLabelPolicy(ctx context.Context, defaultPolicy, preview bool, queries *query.Queries) (*query.LabelPolicy, error) {
	if defaultPolicy {
		if preview {
			return queries.DefaultPreviewLabelPolicy(ctx)
		}
		return queries.DefaultActiveLabelPolicy(ctx)
	}
	if preview {
		return queries.PreviewLabelPolicyByOrg(ctx, authz.GetCtxData(ctx).OrgID)
	}
	return queries.ActiveLabelPolicyByOrg(ctx, authz.GetCtxData(ctx).OrgID)
}

func getLabelPolicyBucketName(ctx context.Context, defaultPolicy, preview bool, queries *query.Queries) string {
	if defaultPolicy {
		return domain.IAMID
	}
	policy, err := getLabelPolicy(ctx, defaultPolicy, preview, queries)
	if err != nil {
		return ""
	}
	if policy.IsDefault {
		return domain.IAMID
	}
	return authz.GetCtxData(ctx).OrgID
}
