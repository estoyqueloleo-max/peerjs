# 📡 Servidor de Señalización PeerJS en Go

Servidor de señalización ligero para WebRTC / PeerJS implementado en Go (`github.com/muka/peerjs-go/server`).

## 🚀 Despliegue en la Nube (Render / PaaS)
El servidor está preparado para plataformas como Render usando `render.yaml`:
```bash
go run main.go
```
* Variable `PORT`: Puerto de escucha (por defecto `9000`).
* Ruta WebSocket / HTTP: `/myapp`

---

## 🏠 Opciones de Autoalojamiento Doméstico (Para usuarios en casa)

Para facilitar que usuarios no técnicos alojen su propio servidor PeerJS y TURN en sus hogares:

### 1. Binario Go "Todo en Uno" (Zero-Config + Pion + UPnP + ACME)
* Un único binario que integra **PeerJS Server** + **Pion TURN** en Go.
* Utiliza **UPnP** para abrir puertos en el router automáticamente.
* Genera certificados SSL/TLS automáticos con Let's Encrypt / DuckDNS vía ACME.
* Muestra un código QR en consola para emparejar la app móvil al instante.

### 2. Túnel Inverso (Cloudflare Tunnel / Tailscale Funnel)
* Permite exponer este servidor Go a internet sin abrir puertos en el router y superando CGNAT (Digi, MásMóvil, 4G/5G).
* Da HTTPS y WebSockets con certificado SSL automático gratuito.

### 3. Apps para NAS y Mini PCs (CasaOS / Umbrel)
* Empaquetado en Docker / App Store local para instalar con un solo clic.

Para una comparativa detallada de conectividad, CGNAT y redes de operadores, consulta la [Guía P2P](../../p2pt/P2P.md).
