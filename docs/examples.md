# Lumo Examples

This document provides comprehensive examples for all Lumo features and commands.

## Basic Usage

### Natural Language Queries

```bash
# Get help with finding files
lumo "How do I find files by content in Linux?"

# Get help with a specific command
lumo "How to use grep with regular expressions?"

# Get explanations for technical concepts
lumo "What's the difference between TCP and UDP?"

# Get help with system administration
lumo "How to check disk space usage?"

# Get help with Git commands
lumo "How to undo the last Git commit?"
```

## Agent Mode

### Basic Agent Tasks

```bash
# Create a backup of documents
lumo auto:create a backup of my documents folder

# Find and analyze large files
lumo agent:find all large files in the current directory and show their types

# Set up a development environment
lumo agent:set up a python virtual environment with flask and sqlalchemy

# Monitor system resources
lumo agent:monitor CPU and memory usage every 5 seconds

# Clean up temporary files
lumo agent:find and remove all temporary files older than 7 days
```

### Agent Mode REPL Commands

When in the Agent Mode REPL interface:

```
# Execute the current plan
run

# Modify the plan using natural language
refine add a step to compress the backup file

# Add a new step to the plan
add tar -czf backup.tar.gz backup/

# Edit a specific step in the plan
edit 3

# Remove a step from the plan
delete 2

# Reorder steps in the plan
move 4 2

# Show available commands
help

# Exit without executing
exit
```

## Chat Mode

### Starting Chat Mode

```bash
# Start interactive chat mode
lumo chat

# Ask a direct question in chat format
lumo chat:Tell me about Linux file permissions
lumo talk:Tell me about Linux file permissions

# Ask a specific question without entering chat mode
lumo ask:What is the capital of France?
lumo ai:What is the capital of France?
```

### Chat Mode Commands

In the interactive Chat Mode:

```
# Send a message to the AI
How does public key encryption work?

# Show help information
help

# Clear conversation history
clear

# Display conversation history
history

# Start a new conversation
new

# List all conversations
list

# Switch to another conversation
switch 2

# Delete a conversation
delete 1

# Exit chat mode
exit
```

## System Commands

### Shell Commands

```bash
# Execute a shell command (MUST use shell: prefix)
lumo shell:ls -la

# Execute another shell command
lumo shell:find . -name "*.go" -type f

# Note: Shell commands are ONLY executed when explicitly prefixed with "shell:"
# The following will NOT execute as a shell command, but will be processed as an AI query:
lumo ls -la
```

### System Health

```bash
# Get a basic health report
lumo health

# Get a detailed system report
lumo system

# Check specific system components
lumo health:memory
lumo health:disk
lumo health:cpu
lumo health:network

# Alternative syntax for health commands
lumo syshealth:memory
lumo syshealth:disk

# Generate a report with specific focus
lumo report:performance
lumo report:security
lumo report:storage

# Alternative syntax for report commands
lumo sysreport:performance
lumo sysreport:security
```

## Internet Speed Testing

```bash
# Run a complete speed test
lumo speed

# Test only download speed
lumo speed:download

# Test only upload speed
lumo speed:upload

# Alternative syntax for speed tests
lumo speedtest:download
lumo speed-test:upload

# Test with natural language
lumo "check my internet speed"
lumo "how fast is my internet connection"
```

## Clipboard Operations

```bash
# Show current clipboard contents
lumo clipboard

# Copy text to clipboard
lumo clipboard "Hello World"

# Append text to existing clipboard content
lumo clipboard append "More text"

# Clear clipboard contents
lumo clipboard clear

# Copy piped content to clipboard
echo "This is some text" | lumo clipboard

# Append piped content to clipboard
cat file.txt | lumo clipboard append
```

## Project Creation

```bash
# Create a Flutter project with BLoC architecture
lumo create:"Flutter app with bloc architecture"

# Create a Flutter project with Provider state management
lumo create:"Flutter app with provider state management"

# Create a Flutter project with Riverpod
lumo create:"Flutter app with riverpod state management"

# Create a Next.js project with Redux
lumo create:"Next.js app with Redux state management"

# Create a Next.js project with Context API
lumo create:"Next.js project with Context API"

# Create a Next.js project with Zustand
lumo create:"Next.js application using Zustand for state"

# Create a basic Next.js project without specific state management
lumo create:"Simple Next.js project"

# Create a React project with Redux
lumo create:"React app with Redux state management"

# Create a React project with Context API
lumo create:"React project with Context API"

# Create a React project with MobX
lumo create:"React application using MobX for state"

# Create a React project with Recoil
lumo create:"React app with Recoil state management"

# Create a basic React project without specific state management
lumo create:"Simple React project"

# Create a FastAPI project
lumo create:"FastAPI project with SQLAlchemy"

# Create a FastAPI project with specific options
lumo create:"Create a FastAPI REST API for a blog"

# Create a Flask project
lumo create:"Flask web application"

# Create a Flask project with specific options
lumo create:"Create a Flask app with SQLAlchemy and authentication"

# Show help for the create command
lumo create
```

## Desktop Assistant

The desktop assistant allows you to control your desktop environment using natural language commands. It uses AI to understand complex commands and execute them.

```bash
# Close a specific window
lumo desktop:"close firefox window"

# Minimize a window
lumo desktop:"minimize terminal window"

# Maximize a window
lumo desktop:"maximize chrome window"

# List all open windows
lumo desktop:"list windows"

# Launch an application
lumo desktop:"launch terminal"

# List running applications
lumo desktop:"list applications"

# Lock the screen
lumo desktop:"lock screen"

# Send a notification
lumo desktop:"send notification Hello World with body This is a test"

# Control media playback
lumo desktop:"play media"
lumo desktop:"pause media"
lumo desktop:"stop media"
lumo desktop:"next track"
lumo desktop:"previous track"

# Change appearance settings (GNOME)
lumo desktop:"set dark mode on"
lumo desktop:"set light mode"
lumo desktop:"change desktop background to /path/to/image.jpg"
lumo desktop:"set GTK theme to Adwaita-dark"
lumo desktop:"change icon theme to Papirus"
lumo desktop:"get current theme"
lumo desktop:"show desktop background"

# Control sound settings (GNOME)
lumo desktop:"set volume to 50 percent"
lumo desktop:"increase volume to 75 percent"
lumo desktop:"mute the sound"
lumo desktop:"unmute the sound"
lumo desktop:"set microphone volume to 80 percent"
lumo desktop:"mute the microphone"
lumo desktop:"show all sound devices"
lumo desktop:"get current volume level"
lumo desktop:"set default sound device to alsa_output.pci-0000_00_1f.3.analog-stereo"

# Control connectivity settings (GNOME)
lumo desktop:"show all network devices"
lumo desktop:"turn on WiFi"
lumo desktop:"turn off WiFi"
lumo desktop:"check WiFi status"
lumo desktop:"enable Bluetooth"
lumo desktop:"disable Bluetooth"
lumo desktop:"check Bluetooth status"
lumo desktop:"turn on airplane mode"
lumo desktop:"turn off airplane mode"
lumo desktop:"check airplane mode status"
lumo desktop:"create a WiFi hotspot with name 'MyHotspot'"
lumo desktop:"create a WiFi hotspot with name 'MyHotspot' and password 'securepass'"
lumo desktop:"turn off WiFi hotspot"
lumo desktop:"check hotspot status"

# AI-powered natural language commands
lumo desktop:"I want to close all Firefox windows and then open a new terminal"
lumo desktop:"Could you please minimize all my windows and then lock my screen?"
lumo desktop:"First open Firefox, then maximize it, and finally play some music"
lumo desktop:"Switch to dark mode and set my background to night-sky.jpg"
lumo desktop:"Increase the volume to 80 percent and then play some music"
lumo desktop:"Turn off WiFi and Bluetooth to save battery"
lumo desktop:"Enable airplane mode and then set volume to 0"
lumo desktop:"Create a hotspot named 'LumoShare' with password 'lumo1234'"
```



## Magic Commands

```bash
# Show a fun dance animation
lumo magic:dance

# List available magic commands
lumo magic:help
```

## Configuration Commands

```bash
# List available AI providers
lumo config:provider list

# Show current AI provider
lumo config:provider show

# Set AI provider
lumo config:provider set gemini
lumo config:provider set openai
lumo config:provider set ollama
lumo config:provider set claude
lumo config:provider set openai-compatible

# List available models for the current provider
lumo config:model list

# Show current model
lumo config:model show

# Set model for current provider (any model string is accepted)
lumo config:model set gemini-2.5-flash-lite
lumo config:model set claude-sonnet-4-6

# Show API key status
lumo config:key show

# Set API key for a provider
lumo config:key set gemini YOUR_API_KEY
lumo config:key set openai YOUR_API_KEY
lumo config:key set claude YOUR_API_KEY
lumo config:key set openai-compatible YOUR_API_KEY

# Remove API key for a provider
lumo config:key remove gemini

# Show current Ollama URL
lumo config:ollama show

# Set Ollama URL
lumo config:ollama set http://localhost:11434

# Test connection to Ollama server
lumo config:ollama test

# Configure an OpenAI-compatible endpoint (xAI, DeepSeek, Mistral, Groq, OpenRouter)
lumo config:compatible show
lumo config:compatible set-url https://api.x.ai/v1
```

## Pipe Support

```bash
# Analyze directory contents
ls -la | lumo

# Explain error logs
cat error.log | lumo

# Understand complex command output
ps aux | grep python | lumo

# Analyze system information
dmesg | lumo

# Get help with command output
ifconfig | lumo

# Analyze JSON data
cat data.json | lumo

# Analyze CSV data
cat data.csv | lumo
```

## File Transfer with Connect

Lumo Connect allows you to transfer files between machines on the same network. For large files (>10MB), it automatically uses chunked transfer for better reliability and performance.

```bash
# Start a server to receive files
lumo connect --receive

# Start a server on a specific port
lumo connect --receive --port 9000

# Specify a custom download directory
lumo connect --receive --path ~/Downloads/transfers

# Discover available Lumo Connect services on the network
lumo connect --discover

# Connect to a peer to send/receive files
lumo connect 192.168.1.5

# Connect to a peer on a specific port
lumo connect 192.168.1.5:9000

# Connect to a peer with a custom download directory
lumo connect 192.168.1.5 --path ~/Downloads/transfers

# Connect to a peer and use chunked transfer for all files (better for large files)
lumo connect 192.168.1.5 --chunked

# Connect to a peer with both custom download directory and chunked transfer
lumo connect 192.168.1.5 --path ~/Downloads/transfers --chunked

# Show connect command help
lumo connect --help

# Access the web interface for Connect (when server is running)
# Open a browser and navigate to: http://localhost:7531/connect/
```

## REST Server Commands

```bash
# Start the REST server daemon
lumo server:start

# Stop the REST server daemon
lumo server:stop

# Check if the server is running
lumo server:status

# Show server help
lumo server:help

# Show current server settings
lumo config:server show

# Enable the REST server
lumo config:server enable

# Disable the REST server
lumo config:server disable

# Control server log messages
lumo config:server quiet on
lumo config:server quiet off

# Enable authentication for the REST server
lumo config:server auth enable

# Disable authentication for the REST server
lumo config:server auth disable

# Change the default admin password
lumo config:server auth password

# Default credentials for the web interface and API:
# Username: admin
# Password: lumo
# Important: Change this password immediately after first login!

# Configure server settings in ~/.config/lumo/config.json:
# - "enable_server": true/false - Enable or disable the REST server
# - "server_port": 7531 - Set the port for the REST server
# - "server_quiet_output": true/false - Control server log messages
# - "enable_auth": true/false - Enable or disable authentication
# - "jwt_secret": "your-secret" - Secret key for JWT token generation
# - "token_expiration_hours": 24 - Token expiration time in hours
# - "refresh_expiration_days": 7 - Refresh token expiration time in days
```

### REST API Endpoints

When the server is running, you can interact with Lumo via HTTP:

```bash
# Check server status (no authentication required)
curl http://localhost:7531/api/v1/status

# Simple ping test to check if server is running (no authentication required)
curl http://localhost:7531/ping

# Chunked File Transfer endpoints (no authentication required)

# Initialize a file upload
curl -X POST -H "Content-Type: application/json" \
  -d '{"filename":"video.mkv","file_size":4831838208}' \
  http://localhost:7531/api/v1/connect/upload/init

# Upload a chunk (replace with your actual upload_id and chunk_id)
curl -X POST -H "Content-Type: application/octet-stream" \
  --data-binary @chunk_file.bin \
  "http://localhost:7531/api/v1/connect/upload/chunk?upload_id=abcdef1234567890&chunk_id=0"

# Complete an upload (replace with your actual upload_id)
curl -X POST "http://localhost:7531/api/v1/connect/upload/complete?upload_id=abcdef1234567890"

# Authentication endpoints

# Login to get a JWT token (no authentication required)
# Using default credentials (username: admin, password: lumo)
curl -X POST -H "Content-Type: application/json" \
  -d '{"username":"admin","password":"lumo"}' \
  http://localhost:7531/api/v1/auth/login

# Refresh an expired token (no authentication required)
curl -X POST -H "Content-Type: application/json" \
  -d '{"refresh_token":"your-refresh-token"}' \
  http://localhost:7531/api/v1/auth/refresh

# Change password (requires authentication)
curl -X POST -H "Content-Type: application/json" \
  -H "Authorization: Bearer your-jwt-token" \
  -d '{"current_password":"lumo","new_password":"new-secure-password"}' \
  http://localhost:7531/api/v1/auth/change-password

# API endpoints (all require authentication when auth is enabled)

# Execute a command (AI query) - Basic usage
curl -X POST -H "Content-Type: application/json" \
  -H "Authorization: Bearer your-jwt-token" \
  -d '{"command":"How do I find large files in Linux?"}' \
  http://localhost:7531/api/v1/execute

# Execute a shell command
curl -X POST -H "Content-Type: application/json" \
  -H "Authorization: Bearer your-jwt-token" \
  -d '{"command":"ls -la", "type":"shell"}' \
  http://localhost:7531/api/v1/execute

# Execute an agent command with parameters
curl -X POST -H "Content-Type: application/json" \
  -H "Authorization: Bearer your-jwt-token" \
  -d '{"command":"create a backup", "type":"agent", "params":{"path":"/home/user/docs"}}' \
  http://localhost:7531/api/v1/execute

# Get system health information
curl -X POST -H "Content-Type: application/json" \
  -H "Authorization: Bearer your-jwt-token" \
  -d '{"command":"check system health", "type":"system_health"}' \
  http://localhost:7531/api/v1/execute

# Generate a system report
curl -X POST -H "Content-Type: application/json" \
  -H "Authorization: Bearer your-jwt-token" \
  -d '{"command":"generate report", "type":"system_report"}' \
  http://localhost:7531/api/v1/execute

# Run a speed test
curl -X POST -H "Content-Type: application/json" \
  -H "Authorization: Bearer your-jwt-token" \
  -d '{"command":"test download speed", "type":"speed_test"}' \
  http://localhost:7531/api/v1/execute

# Get help information
curl -X POST -H "Content-Type: application/json" \
  -H "Authorization: Bearer your-jwt-token" \
  -d '{"command":"help", "type":"help"}' \
  http://localhost:7531/api/v1/execute

# Modify configuration
curl -X POST -H "Content-Type: application/json" \
  -H "Authorization: Bearer your-jwt-token" \
  -d '{"command":"provider show", "type":"config"}' \
  http://localhost:7531/api/v1/execute
```

### Using the REST API with Python

Here's an example of how to use the Lumo REST API with Python:

```python
import requests
import json

# Base URL for the Lumo REST API
base_url = "http://localhost:7531"

# Check server status (no authentication required)
response = requests.get(f"{base_url}/api/v1/status")
print("Server Status:", response.json())

# Login to get authentication tokens
# Using default credentials (username: admin, password: lumo)
login_payload = {
    "username": "admin",
    "password": "lumo"  # Replace with your actual password after changing the default
}
login_response = requests.post(
    f"{base_url}/api/v1/auth/login",
    headers={"Content-Type": "application/json"},
    data=json.dumps(login_payload)
)
auth_data = login_response.json()
token = auth_data["token"]
refresh_token = auth_data["refresh_token"]
print(f"Logged in as: {auth_data['username']}")

# Execute an AI query with authentication
payload = {
    "command": "What is the capital of France?"
}
response = requests.post(
    f"{base_url}/api/v1/execute",
    headers={
        "Content-Type": "application/json",
        "Authorization": f"Bearer {token}"
    },
    data=json.dumps(payload)
)
print("AI Response:", response.json()["output"])

# Execute a shell command with authentication
payload = {
    "command": "df -h",
    "type": "shell"
}
response = requests.post(
    f"{base_url}/api/v1/execute",
    headers={
        "Content-Type": "application/json",
        "Authorization": f"Bearer {token}"
    },
    data=json.dumps(payload)
)
print("Shell Command Output:", response.json()["output"])

# Refresh token when it expires
def refresh_auth_token():
    refresh_payload = {
        "refresh_token": refresh_token
    }
    refresh_response = requests.post(
        f"{base_url}/api/v1/auth/refresh",
        headers={"Content-Type": "application/json"},
        data=json.dumps(refresh_payload)
    )
    new_auth_data = refresh_response.json()
    return new_auth_data["token"], new_auth_data["refresh_token"]

# Change password
def change_password(current_password, new_password, token):
    payload = {
        "current_password": current_password,
        "new_password": new_password
    }
    response = requests.post(
        f"{base_url}/api/v1/auth/change-password",
        headers={
            "Content-Type": "application/json",
            "Authorization": f"Bearer {token}"
        },
        data=json.dumps(payload)
    )
    return response.json()

# Example of chunked file upload with Python
def upload_large_file(file_path, server_url="http://localhost:7531"):
    """Upload a large file using chunked transfer."""
    import os
    import requests
    import json

    # Get file info
    file_size = os.path.getsize(file_path)
    file_name = os.path.basename(file_path)

    print(f"Uploading {file_name} ({file_size} bytes)")

    # Initialize upload
    init_response = requests.post(
        f"{server_url}/api/v1/connect/upload/init",
        headers={"Content-Type": "application/json"},
        data=json.dumps({
            "filename": file_name,
            "file_size": file_size
        })
    )

    if not init_response.ok:
        raise Exception(f"Failed to initialize upload: {init_response.text}")

    upload_data = init_response.json()
    upload_id = upload_data["upload_id"]
    chunk_size = upload_data["chunk_size"]
    total_chunks = len(upload_data["chunks"])

    print(f"Upload initialized with ID: {upload_id}")
    print(f"Chunk size: {chunk_size} bytes")
    print(f"Total chunks: {total_chunks}")

    # Upload chunks
    with open(file_path, "rb") as f:
        for chunk_id in range(total_chunks):
            # Calculate progress
            progress = (chunk_id + 1) * 100 // total_chunks
            print(f"Uploading chunk {chunk_id+1}/{total_chunks} ({progress}%)")

            # Read chunk
            chunk_data = f.read(chunk_size)

            # Upload chunk
            chunk_response = requests.post(
                f"{server_url}/api/v1/connect/upload/chunk?upload_id={upload_id}&chunk_id={chunk_id}",
                headers={"Content-Type": "application/octet-stream"},
                data=chunk_data
            )

            if not chunk_response.ok:
                raise Exception(f"Failed to upload chunk {chunk_id}: {chunk_response.text}")

    # Complete upload
    complete_response = requests.post(
        f"{server_url}/api/v1/connect/upload/complete?upload_id={upload_id}"
    )

    if not complete_response.ok:
        raise Exception(f"Failed to complete upload: {complete_response.text}")

    result = complete_response.json()
    print(f"Upload completed successfully!")
    print(f"File saved to: {result['file_path']}")
    return result["file_path"]

# Example usage:
# upload_large_file("/path/to/large/file.mp4")
```

### Using the REST API with JavaScript/Node.js

Here's an example of how to use the Lumo REST API with JavaScript:

```javascript
const fetch = require('node-fetch');

// Base URL for the Lumo REST API
const baseUrl = 'http://localhost:7531';
let authToken = '';
let refreshToken = '';

// Check server status (no authentication required)
fetch(`${baseUrl}/api/v1/status`)
  .then(response => response.json())
  .then(data => console.log('Server Status:', data))
  .catch(error => console.error('Error:', error));

// Login to get authentication tokens
// Using default credentials (username: admin, password: lumo)
const loginCredentials = {
  username: 'admin',
  password: 'lumo'  // Replace with your actual password after changing the default
};

fetch(`${baseUrl}/api/v1/auth/login`, {
  method: 'POST',
  headers: {
    'Content-Type': 'application/json'
  },
  body: JSON.stringify(loginCredentials)
})
  .then(response => response.json())
  .then(data => {
    authToken = data.token;
    refreshToken = data.refresh_token;
    console.log(`Logged in as: ${data.username}`);

    // Now that we have the token, we can make authenticated requests
    executeAIQuery();
    executeShellCommand();
  })
  .catch(error => console.error('Login Error:', error));

// Execute an AI query with authentication
function executeAIQuery() {
  const aiQuery = {
    command: 'What is the capital of France?'
  };

  fetch(`${baseUrl}/api/v1/execute`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
      'Authorization': `Bearer ${authToken}`
    },
    body: JSON.stringify(aiQuery)
  })
    .then(response => {
      if (response.status === 401) {
        // Token might be expired, try to refresh
        return refreshAuthToken().then(() => executeAIQuery());
      }
      return response.json();
    })
    .then(data => console.log('AI Response:', data.output))
    .catch(error => console.error('Error:', error));
}

// Execute a shell command with authentication
function executeShellCommand() {
  const shellCommand = {
    command: 'df -h',
    type: 'shell'
  };

  fetch(`${baseUrl}/api/v1/execute`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
      'Authorization': `Bearer ${authToken}`
    },
    body: JSON.stringify(shellCommand)
  })
    .then(response => {
      if (response.status === 401) {
        // Token might be expired, try to refresh
        return refreshAuthToken().then(() => executeShellCommand());
      }
      return response.json();
    })
    .then(data => console.log('Shell Command Output:', data.output))
    .catch(error => console.error('Error:', error));
}

// Refresh token when it expires
function refreshAuthToken() {
  const refreshPayload = {
    refresh_token: refreshToken
  };

  return fetch(`${baseUrl}/api/v1/auth/refresh`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify(refreshPayload)
  })
    .then(response => response.json())
    .then(data => {
      authToken = data.token;
      refreshToken = data.refresh_token;
      console.log('Token refreshed successfully');
      return data;
    })
    .catch(error => {
      console.error('Token refresh error:', error);
      throw error;
    });
}

// Change password
function changePassword(currentPassword, newPassword) {
  const payload = {
    current_password: currentPassword,
    new_password: newPassword
  };

  return fetch(`${baseUrl}/api/v1/auth/change-password`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
      'Authorization': `Bearer ${authToken}`
    },
    body: JSON.stringify(payload)
  })
    .then(response => response.json())
    .then(data => {
      console.log('Password changed successfully');
      return data;
    })
    .catch(error => {
      console.error('Password change error:', error);
      throw error;
    });
}

// Example of chunked file upload with JavaScript
async function uploadLargeFile(file) {
  const baseUrl = 'http://localhost:7531';

  console.log(`Uploading ${file.name} (${file.size} bytes)`);

  try {
    // Initialize upload
    const initResponse = await fetch(`${baseUrl}/api/v1/connect/upload/init`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({
        filename: file.name,
        file_size: file.size
      })
    });

    if (!initResponse.ok) {
      throw new Error(`Failed to initialize upload: ${await initResponse.text()}`);
    }

    const uploadData = await initResponse.json();
    const uploadId = uploadData.upload_id;
    const chunkSize = uploadData.chunk_size;
    const totalChunks = uploadData.chunks.length;

    console.log(`Upload initialized with ID: ${uploadId}`);
    console.log(`Chunk size: ${chunkSize} bytes`);
    console.log(`Total chunks: ${totalChunks}`);

    // Upload chunks
    for (let chunkId = 0; chunkId < totalChunks; chunkId++) {
      // Calculate progress
      const progress = Math.floor((chunkId + 1) * 100 / totalChunks);
      console.log(`Uploading chunk ${chunkId+1}/${totalChunks} (${progress}%)`);

      // Calculate chunk boundaries
      const start = chunkId * chunkSize;
      const end = Math.min(start + chunkSize, file.size);

      // Read chunk
      const chunk = file.slice(start, end);

      // Upload chunk
      const chunkResponse = await fetch(
        `${baseUrl}/api/v1/connect/upload/chunk?upload_id=${uploadId}&chunk_id=${chunkId}`,
        {
          method: 'POST',
          headers: {
            'Content-Type': 'application/octet-stream'
          },
          body: chunk
        }
      );

      if (!chunkResponse.ok) {
        throw new Error(`Failed to upload chunk ${chunkId}: ${await chunkResponse.text()}`);
      }
    }

    // Complete upload
    const completeResponse = await fetch(
      `${baseUrl}/api/v1/connect/upload/complete?upload_id=${uploadId}`,
      {
        method: 'POST'
      }
    );

    if (!completeResponse.ok) {
      throw new Error(`Failed to complete upload: ${await completeResponse.text()}`);
    }

    const result = await completeResponse.json();
    console.log(`Upload completed successfully!`);
    console.log(`File saved to: ${result.file_path}`);
    return result.file_path;
  } catch (error) {
    console.error('Upload error:', error);
    throw error;
  }
}

// Example usage:
// const fileInput = document.getElementById('fileInput');
// fileInput.addEventListener('change', async (event) => {
//   const file = event.target.files[0];
//   if (file) {
//     try {
//       const filePath = await uploadLargeFile(file);
//       console.log(`File uploaded to: ${filePath}`);
//     } catch (error) {
//       console.error('Upload failed:', error);
//     }
//   }
// });
```

### Using the REST API with HTML/JavaScript (Web Interface)

Here's a simple HTML page that provides a web interface for Lumo:

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Lumo Web Interface</title>
    <style>
        body {
            font-family: Arial, sans-serif;
            max-width: 800px;
            margin: 0 auto;
            padding: 20px;
        }
        .container {
            display: flex;
            flex-direction: column;
            gap: 10px;
        }
        textarea {
            width: 100%;
            height: 100px;
            padding: 10px;
            border-radius: 5px;
            border: 1px solid #ccc;
        }
        .input-row {
            display: flex;
            gap: 10px;
        }
        input, select, button {
            padding: 10px;
            border-radius: 5px;
            border: 1px solid #ccc;
        }
        button {
            background-color: #4CAF50;
            color: white;
            border: none;
            cursor: pointer;
        }
        button:hover {
            background-color: #45a049;
        }
        pre {
            background-color: #f5f5f5;
            padding: 15px;
            border-radius: 5px;
            overflow-x: auto;
            white-space: pre-wrap;
        }
    </style>
</head>
<body>
    <h1>Lumo Web Interface</h1>
    <div class="container">
        <div class="input-row">
            <input type="text" id="command" placeholder="Enter your command" style="flex-grow: 1;">
            <select id="commandType">
                <option value="">AI (default)</option>
                <option value="shell">Shell</option>
                <option value="agent">Agent</option>
                <option value="system_health">System Health</option>
                <option value="system_report">System Report</option>
                <option value="help">Help</option>
                <option value="config">Config</option>
            </select>
            <button onclick="executeCommand()">Execute</button>
        </div>
        <h3>Response:</h3>
        <pre id="response">Results will appear here...</pre>
    </div>

    <script>
        async function executeCommand() {
            const command = document.getElementById('command').value;
            const commandType = document.getElementById('commandType').value;
            const responseElement = document.getElementById('response');

            if (!command) {
                responseElement.textContent = "Please enter a command";
                return;
            }

            responseElement.textContent = "Processing...";

            try {
                const payload = {
                    command: command
                };

                if (commandType) {
                    payload.type = commandType;
                }

                const response = await fetch('http://localhost:7531/api/v1/execute', {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json'
                    },
                    body: JSON.stringify(payload)
                });

                const data = await response.json();

                if (data.success) {
                    responseElement.textContent = data.output;
                } else {
                    responseElement.textContent = `Error: ${data.error || 'Unknown error'}`;
                }
            } catch (error) {
                responseElement.textContent = `Error: ${error.message}`;
                console.error('Error:', error);
            }
        }
    </script>
</body>
</html>
```

Save this HTML file and open it in a browser while the Lumo server is running to interact with Lumo through a web interface.

## Command-Line Options

```bash
# Display help information
lumo --help
lumo -h
lumo help

# Show version information
lumo --version
lumo -v
lumo version
```
