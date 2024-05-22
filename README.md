# Git Telegram Webhook

Simple Go web server that receives GitHub webhook events and sends a formatted message to a Telegram chat. Designed for
use on [Vercel](https://vercel.com/docs/functions/runtimes/go).

## Supported Webhook Events

- [`push`](https://docs.github.com/en/webhooks/webhook-events-and-payloads#push)

## Setup

1. Fork the repository to your GitHub account.
2. Create a new Vercel project and import the forked repository.
3. Set the required environment variables (see the table below) in the Vercel project settings.

| Environment Variable    | Description                                                                  |
|-------------------------|------------------------------------------------------------------------------|
| `TELEGRAM_BOT_TOKEN`    | The API token for your Telegram bot.                                         |
| `GITHUB_WEBHOOK_SECRET` | The secret key used to verify the integrity of GitHub webhook payloads.      |
| `PORT`                  | The port to listen on. Defaults to `8080`. Only relevant for testing locally |

For debugging purposes, you can set "DANGEROUS_SKIP_GITHUB_WEBHOOK_SECRET_CHECK"
