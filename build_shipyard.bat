@echo off
echo 🏗️ Building the Republic Decentralized Fleet Network...

REM --- Root Config ---
echo. > .env.example
echo. > docker-compose.yml

REM --- Docs ---
mkdir docs 2>nul

REM --- Infrastructure ---
mkdir infrastructure\k8s 2>nul
mkdir infrastructure\docker 2>nul

REM --- Shared Libraries ---
mkdir shared\kafka 2>nul
mkdir shared\schemas 2>nul

REM --- Kotlin: Hull Fabrication (Spring Boot structure) ---
mkdir services\hull-fabrication\src\main\kotlin\com\republic\hull\controllers 2>nul
mkdir services\hull-fabrication\src\main\kotlin\com\republic\hull\models 2>nul
mkdir services\hull-fabrication\src\main\kotlin\com\republic\hull\repositories 2>nul
mkdir services\hull-fabrication\src\main\kotlin\com\republic\hull\services 2>nul
mkdir services\hull-fabrication\src\main\kotlin\com\republic\hull\events 2>nul
mkdir services\hull-fabrication\src\main\resources 2>nul

REM --- Golang: Hyperdrive Calibration ---
mkdir services\hyperdrive-calibration\internal\routes 2>nul
mkdir services\hyperdrive-calibration\internal\calculator 2>nul
mkdir services\hyperdrive-calibration\internal\models 2>nul

REM --- TypeScript: Astromech Provisioning ---
mkdir services\astromech-provisioning\src\controllers 2>nul
mkdir services\astromech-provisioning\src\consumers 2>nul
mkdir services\astromech-provisioning\src\models 2>nul

REM --- Python: Oracle Engine (ML) ---
mkdir services\oracle-engine\app\routers 2>nul
mkdir services\oracle-engine\app\inference 2>nul
mkdir services\oracle-engine\app\streaming 2>nul

REM --- Python: Astromech Oracle LLM (RAG) ---
mkdir services\astromech-oracle-llm\app\routers 2>nul
mkdir services\astromech-oracle-llm\app\rag 2>nul
mkdir services\astromech-oracle-llm\app\vectorstore 2>nul

REM --- ML Training Grounds ---
mkdir ml\notebooks 2>nul
mkdir ml\training 2>nul
mkdir ml\models 2>nul
mkdir ml\data\raw 2>nul
mkdir ml\data\republic_manuals 2>nul

REM --- CI/CD ---
mkdir .github\workflows 2>nul

echo.
echo ✅ Shipyard scaffolding complete!
echo May the Force be with your build.
pause
