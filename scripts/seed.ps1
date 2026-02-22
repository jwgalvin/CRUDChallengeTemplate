$baseUrl = "http://localhost:8080"

$items = @(
    # Backend services (golang)
    @{ name = "Go Backend API";           tags = @("golang", "backend", "rest") },
    @{ name = "Auth Service";             tags = @("auth", "golang", "backend", "jwt", "oauth2", "middleware", "security", "rest", "api", "sessions") },
    @{ name = "Notification Service";     tags = @("golang", "backend", "messaging") },
    @{ name = "File Upload Service";      tags = @("golang", "backend", "storage", "s3", "multipart", "streaming", "files", "upload", "rest", "api", "middleware", "auth", "validation", "compression", "images", "video", "audio", "docs", "binary", "cdn") },
    @{ name = "Email Service";            tags = @("golang", "backend", "email") },
    @{ name = "Payment Gateway";          tags = @("golang", "backend", "payments", "stripe", "webhooks", "idempotency", "fraud", "pci", "billing", "subscriptions", "invoicing", "refunds", "disputes", "crypto", "currency", "tax", "reporting", "compliance", "audit", "retry") },
    @{ name = "Search Service";           tags = @("golang", "backend", "search") },
    @{ name = "Rate Limiter";             tags = @("golang", "backend", "middleware") },
    @{ name = "Job Scheduler";            tags = @("golang", "backend", "cron") },
    @{ name = "WebSocket Server";         tags = @("golang", "backend", "realtime") },

    # Backend services (python)
    @{ name = "Python ML Service";        tags = @("python", "ml", "backend") },
    @{ name = "Data Pipeline";            tags = @("python", "etl", "backend") },
    @{ name = "Analytics Service";        tags = @("python", "analytics", "backend") },
    @{ name = "Image Processing API";     tags = @("python", "ml", "images") },
    @{ name = "NLP Service";              tags = @("python", "ml", "nlp") },

    # Frontend
    @{ name = "React Dashboard";          tags = @("react", "frontend", "typescript") },
    @{ name = "Admin Panel";              tags = @("react", "frontend", "typescript") },
    @{ name = "Vue Storefront";           tags = @("vue", "frontend", "javascript") },
    @{ name = "Angular CRM";             tags = @("angular", "frontend", "typescript") },
    @{ name = "Mobile App";              tags = @("react", "frontend", "mobile") },

    # Databases
    @{ name = "Postgres Database";        tags = @("database", "sql", "postgres") },
    @{ name = "MySQL Replica";            tags = @("database", "sql", "mysql") },
    @{ name = "MongoDB Store";            tags = @("database", "nosql", "mongo") },
    @{ name = "Elasticsearch Cluster";    tags = @("database", "search", "elastic") },
    @{ name = "TimescaleDB";              tags = @("database", "sql", "timeseries") },

    # Caching & Messaging
    @{ name = "Redis Cache";              tags = @("cache", "redis", "backend") },
    @{ name = "Memcached Layer";          tags = @("cache", "backend") },
    @{ name = "Kafka Broker";             tags = @("messaging", "kafka", "backend") },
    @{ name = "RabbitMQ Setup";           tags = @("messaging", "rabbitmq", "backend") },
    @{ name = "NATS Server";              tags = @("messaging", "nats", "backend") },

    # DevOps & Infrastructure
    @{ name = "Docker Setup";             tags = @("docker", "devops", "containers") },
    @{ name = "Kubernetes Manifests";     tags = @("kubernetes", "devops", "containers") },
    @{ name = "Terraform Infra";          tags = @("terraform", "devops", "iac") },
    @{ name = "CI/CD Pipeline";           tags = @("github-actions", "devops", "cicd") },
    @{ name = "Prometheus Monitoring";    tags = @("monitoring", "devops", "metrics") },
    @{ name = "Grafana Dashboard";        tags = @("monitoring", "devops", "metrics") },
    @{ name = "Nginx Config";             tags = @("nginx", "devops", "proxy") },
    @{ name = "Vault Secrets";            tags = @("security", "devops", "secrets") },
    @{ name = "Load Balancer";            tags = @("devops", "networking", "proxy") },
    @{ name = "CDN Config";               tags = @("devops", "networking", "cdn") },

    # Tools & Libraries
    @{ name = "CLI Tool";                 tags = @("golang", "cli") },
    @{ name = "SDK Generator";            tags = @("golang", "cli", "codegen") },
    @{ name = "Migration Runner";         tags = @("golang", "database", "cli") },
    @{ name = "Config Loader";            tags = @("golang", "library") },
    @{ name = "Logger Library";           tags = @("golang", "library", "logging") },

    # Security
    @{ name = "OAuth2 Provider";          tags = @("auth", "security", "backend") },
    @{ name = "API Gateway";              tags = @("security", "backend", "proxy") },
    @{ name = "WAF Rules";                tags = @("security", "devops", "networking") },
    @{ name = "Cert Manager";             tags = @("security", "devops", "tls") },
    @{ name = "Audit Logger";             tags = @("security", "golang", "backend") },

    # Testing & Docs
    @{ name = "Load Test Suite";          tags = @("testing", "k6", "devops") },
    @{ name = "Integration Tests";        tags = @("testing", "golang") },
    @{ name = "API Docs";                 tags = @("docs", "openapi", "rest") },
    @{ name = "Postman Collection";       tags = @("docs", "testing", "rest") },
    @{ name = "Architecture Diagrams";    tags = @("docs") }
)

foreach ($item in $items) {
    $body = $item | ConvertTo-Json
    try {
        $res = Invoke-WebRequest -UseBasicParsing -Uri "$baseUrl/items" `
            -Method POST `
            -ContentType "application/json" `
            -Body $body
        $created = $res.Content | ConvertFrom-Json
        Write-Host "Created: $($created.id)  $($created.name)"
    } catch {
        Write-Host "Failed to create '$($item.name)': $($_.ErrorDetails.Message)"
    }
}
