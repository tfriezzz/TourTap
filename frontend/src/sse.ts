let source: EventSource | null = null

export function startSSE(onMessage: (msg: string) => void) {
  if (source) return

  source = new EventSource("http://localhost:8080/events", {
    withCredentials: true
  })

  source.onmessage = (event: MessageEvent<string>) => {
    onMessage(event.data)
  }

  source.onerror = () => {
    console.warn("SSE disconnected")
  }
}

export function stopSSE() {
  source?.close()
  source = null
}
