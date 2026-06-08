# WhatsApp CLI

Envie mensagens do WhatsApp direto do terminal usando o protocolo Multi-Device.

## Requisitos

- Go 1.21+
- Uma conta no WhatsApp (número de telefone)

## Instalação

```bash
git clone <repo-url> whatsapp-cli
cd whatsapp-cli
go build -o whatsapp-cli .
```

## Uso

### 1. Login (QR Code)

```bash
./whatsapp-cli login
```

Um QR Code será exibido no terminal. Escaneie com seu celular:

1. Abra o WhatsApp no seu celular
2. Toque em **⋮** (Android) ou **Configurações** (iPhone)
3. **Dispositivos Conectados** → **Conectar um Dispositivo**
4. Escaneie o QR Code

A sessão é salva localmente (`session/whatsapp.db`). Você só precisa escanear uma vez.

### 2. Enviar mensagem

```bash
./whatsapp-cli send 5511999999999 "Olá do terminal!"
```

Informe o número no formato internacional **sem** `+` ou espaços.

### 3. Logout

```bash
./whatsapp-cli logout
```

## Exemplo completo

```bash
go build -o whatsapp-cli .
./whatsapp-cli login
# Escaneie o QR com o celular e pressione Ctrl+C

./whatsapp-cli send 5511988887777 "Bom dia!"
# Message sent! ID: 3EB08123DB312947785C27
```

## Como funciona

- Usa a biblioteca [whatsmeow](https://github.com/tulir/whatsmeow) (API Multi-Device)
- Sessão persistida em SQLite via `modernc.org/sqlite` (sem CGO)
- QR Code exibido no terminal com `qrterminal`
