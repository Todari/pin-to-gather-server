package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"

	"github.com/Todari/pin-to-gather-server/models"
)

type LocalSearchService struct {
	NaverClientID     string
	NaverClientSecret string
}

func NewLocalSearchService() *LocalSearchService {
	return &LocalSearchService{
		NaverClientID:     os.Getenv("NAVER_CLIENT_ID"),
		NaverClientSecret: os.Getenv("NAVER_CLIENT_SECRET"),
	}
}

func (s *LocalSearchService) SearchLocal(query string) (models.NaverLocalSearchResponse, error) {
	baseURL := "https://openapi.naver.com/v1/search/local.json"
	params := url.Values{}
	params.Add("query", query)
	params.Add("display", "5")

	req, err := http.NewRequest("GET", baseURL+"?"+params.Encode(), nil)
	if err != nil {
		return models.NaverLocalSearchResponse{}, fmt.Errorf("요청 생성 실패: %w", err)
	}

	req.Header.Set("X-Naver-Client-Id", s.NaverClientID)
	req.Header.Set("X-Naver-Client-Secret", s.NaverClientSecret)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return models.NaverLocalSearchResponse{}, fmt.Errorf("API 호출 실패: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return models.NaverLocalSearchResponse{}, fmt.Errorf("응답 읽기 실패: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return models.NaverLocalSearchResponse{}, fmt.Errorf("네이버 API 오류: %s", string(body))
	}

	var searchResult models.NaverLocalSearchResponse
	if err := json.Unmarshal(body, &searchResult); err != nil {
		return models.NaverLocalSearchResponse{}, fmt.Errorf("JSON 파싱 실패: %w", err)
	}

	return searchResult, nil
}
