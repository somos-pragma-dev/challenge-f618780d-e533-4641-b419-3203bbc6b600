# Prompt para Mejorar el Codigo Base

Copia y pega el siguiente contenido completo en un asistente de IA (Claude, ChatGPT, etc.)
para obtener un ZIP con el proyecto arrancable. Si el adjunto es una carcasa (docs/placeholders),
el asistente debe materializar la estructura del stack del briefing, sin resolver las fases del reto.

---

```
## Briefing del reto (autoridad)
Este bloque manda sobre los archivos adjuntos. El stack y el rol salen de AQUÍ, no de un topic genérico ni de markdown placeholder.

### Contexto técnico original
Diseñar microservicios en Go con gRPC, Docker y Kubernetes

### Reto
- Tema: Go Microservices
- Seniority: junior-l2
- Tipo: mixed
- Título: Diseño y Despliegue de Microservicios en Go
- Tiempo estimado: 12 horas

### Fases (trabajo del HUMANO — PROHIBIDO completarlas)
No implementes estos entregables. Dejalos como hueco pedagógico. El asistente solo materializa el proyecto arrancable para que el participante pueda trabajar.
- Fase 1: Exploración del Dominio — objetivo: Entender las interacciones y restricciones del sistema de microservicios en el contexto de banca digital. — entregable (NO resolver): Mapa de interacciones y restricciones del sistema de microservicios.
- Fase 2: Diseño de Microservicios — objetivo: Diseñar los microservicios para garantizar la consistencia de datos y manejar errores de red y timeouts. — entregable (NO resolver): Especificación detallada de los microservicios, incluyendo endpoints gRPC y estrategias de manejo de errores.
- Fase 3: Despliegue en Kubernetes — objetivo: Desplegar los microservicios en un entorno de Kubernetes y verificar su funcionamiento. — entregable (NO resolver): Microservicios desplegados y funcionando en Kubernetes, con verificación de comunicación y manejo de errores.
- Fase 4: Evaluación y Optimización — objetivo: Evaluar el desempeño del sistema y realizar optimizaciones necesarias. — entregable (NO resolver): Informe de evaluación y optimización del sistema de microservicios.

Eres un asistente experto en análisis, corrección y generación de archivos de cualquier tipo:
código fuente, documentación, hojas de cálculo, documentos Word, configuraciones, entre otros.
Voy a enviarte una cadena de texto que contiene uno o más archivos. Cada archivo está delimitado por un marcador con el siguiente formato:
// === ARCHIVO: ruta/del/archivo.extension ===
o también puede aparecer como:
## === ARCHIVO: ruta/del/archivo.extension ===
Lo que sigue al marcador puede ser:

El contenido real del archivo (código, texto, YAML, etc.)
Una descripción en lenguaje natural de lo que debe contener el archivo


TU TAREA
PASO 0 — ¿Esto es un proyecto o una carcasa?
Antes de extraer archivos, leé el Briefing (si está) y diagnosticá el adjunto.

Es CARCASA si ocurre CUALQUIERA de estas:
- No hay manifiesto de dependencias del stack del briefing (manifest.json de VTEX IO / package.json / pom.xml / build.gradle / requirements.txt / go.mod / *.tf / *.csproj, según corresponda)
- Hay un "binario" que en realidad es un comentario ("no puede ser mostrado como texto plano", placeholder .fig/.docx vacío)
- Los markdowns ya completan entregables de fases posteriores ("se implementó fade-in", lista de áreas ya resuelta)

Si es CARCASA:
- MATERIALIZÁ un proyecto que arranca en el stack del briefing (VTEX IO Store Framework, Angular, Terraform, pytest, Nest, etc.). Incluí manifiesto, punto de entrada y capa de interfaz reales.
- NO copies los markdowns de "solución" como si fueran el producto. Son ruido de generación.
- NO resuelvas las fases del briefing (están marcadas PROHIBIDO). Dejá el hueco pedagógico: el flujo existe, las microinteracciones/calidad/infra que el reto pide NO están hechas.
- Después seguí al PASO 5 (ZIP).

Si es un proyecto REAL (manifiesto + código que compila o arranca):
- Seguí PASO 1 en adelante. 🔴 compilación sí. 🟡 pedagógico no.

PASO 1 — Detección y extracción
Identifica todos los archivos presentes en la cadena. Para cada archivo extrae:

Su ruta completa (ej: src/main/java/com/pragma/Service.java)
Su contenido o descripción

PASO 2 — Clasificación por tipo
Clasifica cada archivo en una de estas categorías:
A) Código fuente (Java, Python, TypeScript, JavaScript, Kotlin, etc.)
B) Configuración / documentación (YAML, properties, Markdown, JSON, txt, etc.)
C) Excel (.xlsx, .xls, .csv)
D) Word (.docx, .doc)
E) Otro tipo de archivo binario o especial
PASO 3 — Clasificación de errores en código fuente

Objetivo prioritario: que el proyecto compile. No corrijas flujo de negocio ni lógica funcional.

Antes de modificar cualquier archivo de código fuente, clasifica cada problema encontrado en una de estas dos categorías:
🔴 ERROR DE COMPILACIÓN — corregir siempre
Son errores que impiden que el proyecto arranque, sin valor pedagógico:

Import faltante o incorrecto
Clase, método o variable referenciada que no existe en ningún archivo del proyecto
Error de sintaxis
Anotación con atributos inválidos
Dependencia ausente en pom.xml, package.json, etc.
Archivo referenciado que no existe y debe ser creado con implementación mínima

→ CORREGIR estos errores.
🟡 PROBLEMA FUNCIONAL O DE CALIDAD — preservar siempre
Son problemas que no impiden compilar. Pueden ser intencionales para el aprendizaje:

Clave secreta hardcodeada ("secret", "password123")
API deprecada que funciona pero tiene reemplazo moderno
Lógica de negocio incorrecta o incompleta
Código redundante o de baja legibilidad
Falta de validaciones en flujo de negocio
Patrones de diseño incorrectos pero funcionales
Concurrencia no segura
Configuración funcional pero no óptima

→ PRESERVAR tal cual. No corregir, no mejorar, no comentar.
PASO 4 — Procesamiento según tipo de archivo
Tipo A — Código fuente
Aplica únicamente las correcciones clasificadas como 🔴 ERROR DE COMPILACIÓN.
No alteres ningún elemento clasificado como 🟡 PROBLEMA FUNCIONAL O DE CALIDAD.
Si falta un archivo referenciado, créalo con la implementación mínima necesaria para compilar.
Tipo B — Configuración / documentación
Extrae el contenido tal cual, sin modificaciones salvo errores evidentes de sintaxis
(ej: YAML mal indentado).
Tipo C — Excel (.xlsx)
Si viene con contenido real, genera el archivo respetando ese contenido.
Si viene con descripción en lenguaje natural, genera un archivo Excel funcional con:

Fila de encabezados en negrita con color de fondo distintivo
Columnas con ancho ajustado al contenido
Tipos de dato correctos por columna
Validaciones si la descripción lo indica
Hojas nombradas descriptivamente si hay más de una
Filas de ejemplo si no hay datos reales

Tipo D — Word (.docx)
Si viene con contenido real, genera el archivo respetando ese contenido.
Si viene con descripción en lenguaje natural, genera un documento Word funcional con:

Estilos de título (Título 1, Título 2) para jerarquía de secciones
Fuente legible (Calibri o equivalente), tamaño 11-12pt para cuerpo
Márgenes estándar
Tabla de contenido si tiene múltiples secciones
Tablas con encabezados en negrita si aplica

Tipo E — Otro
Genera el archivo con el contenido o estructura más apropiada según la descripción.
PASO 5 — Exportación en ZIP
Empaqueta todos los archivos en un único archivo ZIP descargable respetando exactamente
la estructura de rutas indicada por los marcadores.
El ZIP debe incluir:

Archivos de código con únicamente los errores de compilación corregidos
Archivos de configuración y documentación sin cambios
Archivos nuevos creados para resolver dependencias de compilación faltantes
Archivos Excel y Word generados desde descripción

IMPORTANTE: El ZIP debe estar listo para descargar al finalizar. No preguntes si el usuario
quiere generarlo. Simplemente genera el archivo y proporciona el enlace de descarga; No debes desplegar en el chat el resumen de lo que arreglaste al Zip, solo entregalo.

REGLAS IMPORTANTES

No omitas ningún archivo aunque no tenga errores ni modificaciones
Respeta los nombres y rutas exactas indicadas por los marcadores
Si un archivo no tiene marcador claro, infiere el nombre desde su contenido
Si la cadena contiene solo documentación, placeholders o binarios fake, NO la reproduzcas:
aplicá PASO 0 (materializar el proyecto del briefing). Reproducir la carcasa es un fallo.
No agregues texto después del enlace de descarga del ZIP
No preguntes si el usuario quiere el ZIP: simplemente generalo siempre
Si detectas que falta un archivo de configuración necesario para compilar
(pom.xml, package.json, requirements.txt, build.gradle, etc.), créalo e inclúyelo
inferiendo su contenido desde los imports y frameworks detectados en el código
Nunca corrijas problemas 🟡 aunque parezcan obvios o fáciles de mejorar.
El participante que recibirá este proyecto los debe encontrar y resolver él mismo.


INPUT
Aquí está la cadena con los archivos:
package main

import (
	"context"
	"log"
	"net/http"

	pb "github.com/your-repo/auth-service/proto"
	"google.golang.org/grpc"
)

type server struct {}

func (s *server) Authenticate(ctx context.Context, req *pb.AuthRequest) (*pb.AuthResponse, error) {
	// Implement authentication logic here
	return &pb.AuthResponse{Token: "sample-token"}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err!= nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterAuthServiceServer(s, &server{})
	log.Println("Auth service started at :50051")
	if err := s.Serve(lis); err!= nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
// === ARCHIVO: auth-service/main.go ===
package main

import (
	"context"
	"log"

	pb "github.com/your-repo/transaction-service/proto"
	"google.golang.org/grpc"
)

type server struct {}

func (s *server) ProcessTransaction(ctx context.Context, req *pb.TransactionRequest) (*pb.TransactionResponse, error) {
	// Implement transaction processing logic here
	return &pb.TransactionResponse{Status: "success"}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50052")
	if err!= nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterTransactionServiceServer(s, &server{})
	log.Println("Transaction service started at :50052")
	if err := s.Serve(lis); err!= nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
// === ARCHIVO: transaction-service/main.go ===
package main

import (
	"context"
	"log"

	pb "github.com/your-repo/notification-service/proto"
	"google.golang.org/grpc"
)

type server struct {}

func (s *server) SendNotification(ctx context.Context, req *pb.NotificationRequest) (*pb.NotificationResponse, error) {
	// Implement notification sending logic here
	return &pb.NotificationResponse{Status: "sent"}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50053")
	if err!= nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterNotificationServiceServer(s, &server{})
	log.Println("Notification service started at :50053")
	if err := s.Serve(lis); err!= nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
// === ARCHIVO: notification-service/main.go ===
apiVersion: apps/v1
kind: Deployment
metadata:
  name: auth-service
spec:
  replicas: 1
  selector:
    matchLabels:
      app: auth-service
  template:
    metadata:
      labels:
        app: auth-service
    spec:
      containers:
      - name: auth-service
        image: your-repo/auth-service:latest
        ports:
        - containerPort: 50051

---

apiVersion: apps/v1
kind: Deployment
metadata:
  name: transaction-service
spec:
  replicas: 1
  selector:
    matchLabels:
      app: transaction-service
  template:
    metadata:
      labels:
        app: transaction-service
    spec:
      containers:
      - name: transaction-service
        image: your-repo/transaction-service:latest
        ports:
        - containerPort: 50052

---

apiVersion: apps/v1
kind: Deployment
metadata:
  name: notification-service
spec:
  replicas: 1
  selector:
    matchLabels:
      app: notification-service
  template:
    metadata:
      labels:
        app: notification-service
    spec:
      containers:
      - name: notification-service
        image: your-repo/notification-service:latest
        ports:
        - containerPort: 50053
// === ARCHIVO: k8s-config/deployment.yaml ===
apiVersion: v1
kind: Service
metadata:
  name: auth-service
spec:
  selector:
    app: auth-service
  ports:
    - protocol: TCP
      port: 50051
      targetPort: 50051

---

apiVersion: v1
kind: Service
metadata:
  name: transaction-service
spec:
  selector:
    app: transaction-service
  ports:
    - protocol: TCP
      port: 50052
      targetPort: 50052

---

apiVersion: v1
kind: Service
metadata:
  name: notification-service
spec:
  selector:
    app: notification-service
  ports:
    - protocol: TCP
      port: 50053
      targetPort: 50053
// === ARCHIVO: k8s-config/service.yaml ===
FROM golang:1.20 AS builder
WORKDIR /app
COPY..
RUN go build -o main

FROM alpine:latest
RUN apk add --no-cache ca-certificates
COPY --from=builder /app/main /app/
WORKDIR /app
EXPOSE 50051
CMD ["/app/main"]
// === ARCHIVO: Dockerfile ===
version: '3.8'

services:
  auth-service:
    build:
      context:./auth-service
    ports:
      - '50051:50051'
  transaction-service:
    build:
      context:./transaction-service
    ports:
      - '50052:50052'
  notification-service:
    build:
      context:./notification-service
    ports:
      - '50053:50053'
// === ARCHIVO: docker-compose.yml ===
#.dockerignore
/dist
/tmp

# Node modules
/node_modules

# Go build artifacts
*.exe
*.test

# Dependency directories
/Godeps/
/vendor/

# Ignore.dockerignore
.dockerignore
// === ARCHIVO:.dockerignore ===
#!/bin/bash

# Health check script for microservices

# Check if the microservice is running
if curl --output /dev/null --silent --head --fail http://localhost:50051; then
  echo "Microservice is healthy"
  exit 0
else
  echo "Microservice is not healthy"
  exit 1
fi
// === ARCHIVO: healthcheck.sh ===
```
