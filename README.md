# 핀투게더 서버 (Pin-to-Gather Server)

핀투게더 서버는 사용자가 함께 지도를 공유하며 실시간으로 서로가 보고 있는 구역과 지역들을 참조하여 여행이나 만남 일정을 정할 수 있는 서비스의 백엔드입니다.

## 기술 스택

- **Backend**: Go
  - Web Framework: [Gin](https://github.com/gin-gonic/gin)
  - WebSocket: [Gorilla WebSocket](https://github.com/gorilla/websocket)
  - Database: PostgreSQL
  - ORM: [GORM](https://gorm.io/)

- **Containerization**: Docker, Docker Compose

## 주요 기능

- **실시간 지도 공유**: WebSocket을 통해 사용자가 현재 보고 있는 지도의 구역을 다른 사용자와 실시간으로 공유합니다.
- **사용자 관리**: 사용자 연결 및 해제를 관리하고, 실시간으로 사용자 정보를 브로드캐스트합니다.
- **데이터베이스 관리**: PostgreSQL을 사용하여 사용자 및 지도 데이터를 저장하고 관리합니다.

## 설치 및 실행

### 요구 사항

- Docker 및 Docker Compose가 설치되어 있어야 합니다.

### 실행 방법

1. **환경 변수 설정**: `.env` 파일을 생성하고 필요한 환경 변수를 설정합니다.

2. **Docker Compose 실행**: 다음 명령어를 사용하여 서비스를 빌드하고 실행합니다.

   ```bash
   docker-compose up --build
   ```

3. **서버 접속**: 서버는 기본적으로 `http://localhost:8080`에서 실행됩니다.

## 디렉토리 구조

- `api/`: API 핸들러
- `services/`: 비즈니스 로직 및 WebSocket 서비스
- `repository/`: 데이터베이스 접근 로직
- `models/`: 데이터베이스 모델
- `config/`: 설정 파일
- `database/`: 데이터베이스 연결 설정
- `routes/`: 라우팅 설정

## 기여 방법

1. 이슈를 생성합니다.
2. 이 저장소를 포크합니다.
3. 새로운 브랜치를 생성합니다. (`git checkout -b feature/#{이슈번호}`)
4. 변경 사항을 커밋합니다. (`git commit -am 'Add new feature'`)
5. 브랜치에 푸시합니다. (`git push origin feature/#{이슈번호}`)
6. dev 브랜치로 Pull Request를 생성합니다.

## 문의

질문이나 제안 사항이 있으시면 [이슈](https://github.com/Todari/pin-to-gather-server/issues) 페이지를 통해 문의해 주세요.