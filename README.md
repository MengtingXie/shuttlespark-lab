# ShuttleSpark Lab

An advanced platform for badminton match analysis that combines high-fidelity manual tagging with automated computer vision insights. Designed to help players and coaches quantify performance, visualize tactics, and generate data-driven training plans.

## 🚀 Features

* **Interactive Court Logger**: Web-based SVG interface for precise shot placement, movement tracking, and rally recording.
* **Visual Timeline**: Non-linear editor style timeline to visualize rallies, breaks, and key events with adjustable lead/lag times.
* **Tactical Dashboard**: Instant generation of heatmaps, shot distribution charts, and consistency metrics.
* **AI Video Analysis**: (In Progress) Automated shuttlecock tracking, player pose estimation, and shot classification using Computer Vision.
* **Smart Coach**: (Planned) LLM-generated training recommendations based on biomechanical analysis and match statistics.

## 🛠️ Tech Stack

### Frontend

* **React (Vite)**: Optimized for client-side performance and fast interactions.
* **SVG / D3.js**: For interactive court mapping and data visualization.

### Backend (Core)

* **Go (Golang)**: High-performance API Gateway and business logic.
* **Huma**: Modern framework for building APIs with automatic OpenAPI 3.1 & JSON Schema validation.
* **PostgreSQL**: Primary database for match metadata and user data.
* **TimescaleDB**: Optimized storage for high-frequency tracking data (trajectories).

### AI Engine (Microservice)

* **Python**: Environment for ML/CV tasks.
* **FastAPI**: Serving AI models.
* **Celery + Redis**: Asynchronous task queue for handling heavy video processing jobs.
* **Core Models**: TrackNetV3 (Shuttle), YOLOv8 (Player detection), MMPose (Pose estimation).

## 🏗️ Architecture

The system follows a hybrid architecture to leverage the speed of Go for web services and the ecosystem of Python for AI.mermaid
graph TD
User <-->|HTTP/WebSocket| Go
Go <--> DB
Go -->|Enqueue Task| Redis
Redis -->|Consume Task| Python
Python -->|Store Results| DB

```

## 📦 Getting Started

### Prerequisites

- Docker & Docker Compose
- Go 1.22+
- Node.js 18+
- Python 3.10+ (Optional, for local AI dev)

### Quick Start

1.  **Clone the repository**
    ```bash
    git clone [https://github.com/yourusername/badminton-analytics.git](https://github.com/yourusername/badminton-analytics.git)
    ```

2.  **Start Infrastructure (Database & Redis)**
    ```bash
    docker-compose up -d
    ```

3.  **Run Backend**
    ```bash
    cd backend
    go run main.go
    ```

4.  **Run Frontend**
    ```bash
    cd frontend
    npm install
    npm run dev
    ```

## 🗺️ Roadmap

- [x] **Phase 1: Manual Input & Visualization**
    - Interactive SVG court component.
    - Basic match scoring and event logging.
- [ ] **Phase 2: Backend & Persistence**
    - Go API implementation with Huma.
    - PostgreSQL schema design with JSONB support.
- [ ] **Phase 3: AI Pipeline Integration**
    - Video upload and processing workflow.
    - Automated trajectory extraction (TrackNet).
- [ ] **Phase 4: Advanced Analytics**
    - 3D trajectory mapping.
    - AI Coach reporting.

## 🔧 Development Workflow

This project uses [OpenSpec](https://github.com/Fission-AI/OpenSpec) for spec-driven development (SDD) with AI assistants.

### OpenSpec Commands

- `/opsx:new <change-name>` - Start a new feature/change with structured planning
- `/opsx:continue` - Create the next artifact in the workflow
- `/opsx:ff` - Fast-forward through all planning documents
- `/opsx:apply` - Implement the tasks from your plan
- `/opsx:verify` - Verify implementation against specs
- `/opsx:archive` - Archive completed change and update specs

### Workflow

1. **Plan**: Use `/opsx:new` to create a proposal, specs, and design docs
2. **Implement**: Use `/opsx:apply` to execute tasks
3. **Archive**: Use `/opsx:archive` to finalize and update project specs

All changes are tracked in `openspec/changes/` and project specs in `openspec/specs/`.

## 📄 License

MIT