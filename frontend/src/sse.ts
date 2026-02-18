let source: EventSource | null = null

export function startSSE(onMessage: (msg: string) => void) {
  if (source) return

  source = new EventSource("/events", {
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
