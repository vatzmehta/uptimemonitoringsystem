package url

func ToUrlResource(dto UrlResourceDto) UrlResource {
	return UrlResource{
		Url:              dto.Url,
		CrawlTimeout:     dto.CrawlTimeout,
		Frequency:        dto.Frequency,
		FailureThreshold: dto.FailureThreshold,
		Enable:           dto.Enable,
	}
}

func ToUrlResourceDTO(urlObj UrlResource) UrlResourceDto {
	return UrlResourceDto{
		Url:              urlObj.Url,
		CrawlTimeout:     urlObj.CrawlTimeout,
		Frequency:        urlObj.Frequency,
		FailureThreshold: urlObj.FailureThreshold,
		CreatedAt:        urlObj.Model.CreatedAt,
		UpdatedAt:        urlObj.Model.UpdatedAt,
		Enable:           urlObj.Enable,
		Id:               urlObj.Model.ID,
	}

}

func ToUrlResourceDTOs(urlObjs []UrlResource) []UrlResourceDto {
	urlResourceDTOs := make([]UrlResourceDto, len(urlObjs))
	for i, urlObj := range urlObjs {
		urlResourceDTOs[i] = ToUrlResourceDTO(urlObj)
	}
	return urlResourceDTOs
}
