# Diseño y Despliegue de Microservicios en Go

En un entorno de banca digital, se requiere diseñar y desplegar un conjunto de microservicios que interactúen para procesar transacciones financieras. Los microservicios deben comunicarse utilizando gRPC y ser desplegados en un entorno de Kubernetes. El sistema debe manejar un volumen de 10 000 transacciones por minuto con un tiempo de respuesta promedio de 50 ms. Los microservicios incluyen: servicio de autenticación, servicio de procesamiento de transacciones y servicio de notificación. El sistema debe garantizar la consistencia de los datos entre los microservicios y manejar errores de red y timeouts de servicios externos.

## Informacion General

| Campo | Valor |
|-------|-------|
| **Tema** | Go Microservices |
| **Nivel** | junior-l2 |
| **Tipo** | mixed |
| **Tiempo estimado** | 12 horas |

## Fases del Reto

### Fase 0: Configuración del Proyecto

**Objetivo:** Obtener el proyecto base funcional enviando el Código Base a un asistente de IA, que lo analizará, corregirá errores y generará un ZIP listo para usar.

**Tiempo estimado:** 15-30 minutos

**Instrucciones:**

- Asegúrate de tener instalado para ejecutar el proyecto: Un IDE o editor de código.
- Copia todo el contenido del campo **Código Base** de este reto — incluyendo el texto de instrucciones que aparece al inicio.
- Abre un asistente de IA (Claude en claude.ai, ChatGPT o Gemini — se recomienda Claude), pega el contenido copiado en el chat y envíalo.
- El asistente analizará los archivos, corregirá errores y generará un archivo ZIP descargable. Descárgalo y extráelo en la carpeta donde quieras trabajar.
- Verifica que el proyecto arranca sin errores.

**Entregable:** El proyecto compila/arranca sin errores.

<details>
<summary>Pistas de conocimiento</summary>

- Copia el Código Base completo incluyendo el texto de instrucciones al inicio — esas instrucciones le indican al asistente exactamente qué hacer con los archivos.
- Si el asistente no genera el ZIP automáticamente al terminar el análisis, escríbele: "genera el ZIP ahora".
- Si el proyecto tiene errores al arrancar, comparte el mensaje de error con el mismo asistente para que lo corrija.

</details>

### Fase 1: Exploración del Dominio

**Objetivo:** Entender las interacciones y restricciones del sistema de microservicios en el contexto de banca digital.

**Tiempo estimado:** 2 horas

**Instrucciones:**

- Identificar los microservicios clave y sus responsabilidades.
- Enumerar las restricciones y umbrales operativos del sistema.
- Describir las interacciones entre los microservicios y los servicios externos.

**Entregable:** Mapa de interacciones y restricciones del sistema de microservicios.

<details>
<summary>Pistas de conocimiento</summary>

- Considera las propiedades operativas como consistencia, latencia y disponibilidad.
- Piensa en los posibles modos de falla y cómo el sistema debe responder.

</details>

### Fase 2: Diseño de Microservicios

**Objetivo:** Diseñar los microservicios para garantizar la consistencia de datos y manejar errores de red y timeouts.

**Tiempo estimado:** 5 horas

**Instrucciones:**

- Definir la estructura de cada microservicio y sus endpoints gRPC.
- Establecer mecanismos para garantizar la consistencia de datos entre microservicios.
- Implementar estrategias para manejar errores de red y timeouts de servicios externos.

**Entregable:** Especificación detallada de los microservicios, incluyendo endpoints gRPC y estrategias de manejo de errores.

<details>
<summary>Pistas de conocimiento</summary>

- Considera el uso de patrones de diseño como sagas para garantizar la consistencia.
- Piensa en cómo implementar reintentos y backoffs exponenciales para manejar timeouts.

</details>

### Fase 3: Despliegue en Kubernetes

**Objetivo:** Desplegar los microservicios en un entorno de Kubernetes y verificar su funcionamiento.

**Tiempo estimado:** 3 horas

**Instrucciones:**

- Configurar los descriptores de Kubernetes para cada microservicio.
- Desplegar los microservicios en un cluster de Kubernetes.
- Verificar que los microservicios se comuniquen correctamente y manejen los errores definidos.

**Entregable:** Microservicios desplegados y funcionando en Kubernetes, con verificación de comunicación y manejo de errores.

<details>
<summary>Pistas de conocimiento</summary>

- Utiliza herramientas como Helm para simplificar la configuración de Kubernetes.
- Verifica la comunicación entre microservicios y el manejo de errores definidos.

</details>

### Fase 4: Evaluación y Optimización

**Objetivo:** Evaluar el desempeño del sistema y realizar optimizaciones necesarias.

**Tiempo estimado:** 2 horas

**Instrucciones:**

- Medir el tiempo de respuesta y el volumen de transacciones procesadas.
- Identificar cuellos de botella y proponer optimizaciones.
- Implementar y verificar las optimizaciones propuestas.

**Entregable:** Informe de evaluación y optimización del sistema de microservicios.

<details>
<summary>Pistas de conocimiento</summary>

- Utiliza herramientas de monitoreo y profiling para identificar cuellos de botella.
- Propone optimizaciones basadas en los resultados de la evaluación.

</details>

## Dimensiones Evaluadas

- **queEs**: ¿Qué son los microservicios y cómo se comunican en este reto?
- **paraQueSirve**: ¿Para qué sirven los microservicios en el contexto de banca digital?
- **comoSeUsa**: ¿Cómo se usan los microservicios para procesar transacciones financieras?
- **erroresComunes**: ¿Cuáles son los errores comunes que pueden ocurrir en la comunicación entre microservicios y cómo se manejan?
- **queDecisionesImplica**: ¿Qué decisiones implica el diseño y despliegue de microservicios en un entorno de Kubernetes?

## Criterios de Evaluacion

- Identificar y describir los microservicios clave y sus responsabilidades.
- Establecer mecanismos para garantizar la consistencia de datos entre microservicios.
- Implementar estrategias para manejar errores de red y timeouts.
- Desplegar los microservicios en un entorno de Kubernetes y verificar su funcionamiento.
- Evaluar el desempeño del sistema y realizar las optimizaciones necesarias.

## Como trabajar con un asistente de IA

- **AGENTS.md** — instrucciones nativas del repo (Cursor, Codex, Copilot, Gemini, Claude Code). Abrí el proyecto y el agente las carga solo.
- **PROMPT_MEJORA.md** — el mismo prompt, para copiar y pegar en un chat (claude.ai, ChatGPT, etc.).

---

*Reto generado automaticamente por Challenge Generator - Pragma*
