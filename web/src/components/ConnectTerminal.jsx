import { useState, useEffect, useRef } from 'react';

export default function ConnectTerminal({
  isServer = true,
  peerIp = '',
  port = '8080',
  savePath = '~/Downloads',
  showControls = true,
  title = "Lumo Connect"
}) {
  const [isConnected, setIsConnected] = useState(false);
  const [files, setFiles] = useState([]);
  const [logs, setLogs] = useState([]);
  const [dragActive, setDragActive] = useState(false);
  const [analytics, setAnalytics] = useState(null);
  const terminalRef = useRef(null);
  const fileInputRef = useRef(null);

  // Initialize Firebase Analytics
  useEffect(() => {
    const initFirebase = async () => {
      try {
        // Dynamically import Firebase modules to avoid build errors when Firebase is not available
        const { initAnalytics } = await import('../utils/firebase');

        try {
          // Try to import Firebase Analytics
          await import('firebase/analytics');

          // Initialize analytics
          const analyticsInstance = await initAnalytics();
          setAnalytics(analyticsInstance);
        } catch (error) {
          console.log('Firebase Analytics not available, analytics disabled');
        }
      } catch (error) {
        console.log('Firebase module not available, analytics disabled');
      }
    };

    initFirebase();
  }, []);

  // Log events to Firebase Analytics
  const logAnalyticsEvent = async (eventName, eventParams = {}) => {
    if (analytics) {
      try {
        // Dynamically import logEvent to avoid build errors
        const { logEvent } = await import('firebase/analytics');
        logEvent(analytics, eventName, eventParams);
      } catch (error) {
        console.log(`Analytics event not logged: ${eventName}`);
      }
    }
  };

  // Scroll to bottom when logs change
  useEffect(() => {
    if (terminalRef.current) {
      terminalRef.current.scrollTop = terminalRef.current.scrollHeight;
    }
  }, [logs]);

  // Add a log message
  const addLog = (message, type = 'info') => {
    setLogs(prev => [...prev, { message, type, timestamp: new Date() }]);
  };

  // Handle connection
  const handleConnect = () => {
    if (isConnected) {
      setIsConnected(false);
      addLog('Disconnected from peer', 'system');
      logAnalyticsEvent('connect_disconnect', { action: 'disconnect' }).catch(() => {});
    } else {
      setIsConnected(true);
      if (isServer) {
        addLog(`Server started on port ${port}`, 'system');
        addLog(`Waiting for connections...`, 'info');
        logAnalyticsEvent('connect_start_server', { port }).catch(() => {});
      } else {
        addLog(`Connecting to ${peerIp}:${port}...`, 'system');
        setTimeout(() => {
          addLog(`Connected to ${peerIp}:${port}`, 'success');
        }, 1500);
        logAnalyticsEvent('connect_to_peer', { peer_ip: peerIp, port }).catch(() => {});
      }
    }
  };

  // Handle file selection
  const handleFileSelect = () => {
    fileInputRef.current.click();
    logAnalyticsEvent('connect_select_files', { method: 'button' }).catch(() => {});
  };

  // Handle file input change
  const handleFileInputChange = (e) => {
    const selectedFiles = Array.from(e.target.files);
    if (selectedFiles.length > 0) {
      addSelectedFiles(selectedFiles);
    }
  };

  // Add selected files to the list
  const addSelectedFiles = (selectedFiles) => {
    const newFiles = selectedFiles.map(file => ({
      id: Math.random().toString(36).substr(2, 9),
      name: file.name,
      size: file.size,
      type: file.type,
      status: 'queued'
    }));

    setFiles(prev => [...prev, ...newFiles]);

    if (isConnected) {
      // Simulate sending files
      newFiles.forEach(file => {
        addLog(`Sending ${file.name} (${formatFileSize(file.size)})`, 'info');
        simulateFileTransfer(file.id);
      });
    } else {
      addLog('Connect to a peer to send files', 'warning');
    }

    logAnalyticsEvent('connect_files_added', {
      count: selectedFiles.length,
      total_size: selectedFiles.reduce((sum, file) => sum + file.size, 0)
    }).catch(() => {});
  };

  // Simulate file transfer
  const simulateFileTransfer = (fileId) => {
    // Update status to sending
    setFiles(prev => prev.map(file =>
      file.id === fileId ? { ...file, status: 'sending' } : file
    ));

    // Simulate transfer completion after random time (1-3 seconds)
    const transferTime = Math.random() * 2000 + 1000;
    setTimeout(() => {
      setFiles(prev => prev.map(file =>
        file.id === fileId ? { ...file, status: 'sent' } : file
      ));

      const file = files.find(f => f.id === fileId);
      if (file) {
        addLog(`Sent ${file.name} successfully`, 'success');
      }
    }, transferTime);
  };

  // Format file size
  const formatFileSize = (bytes) => {
    if (bytes === 0) return '0 Bytes';
    const k = 1024;
    const sizes = ['Bytes', 'KB', 'MB', 'GB', 'TB'];
    const i = Math.floor(Math.log(bytes) / Math.log(k));
    return parseFloat((bytes / Math.pow(k, i)).toFixed(2)) + ' ' + sizes[i];
  };

  // Handle drag events
  const handleDrag = (e) => {
    e.preventDefault();
    e.stopPropagation();

    if (e.type === 'dragenter' || e.type === 'dragover') {
      setDragActive(true);
    } else if (e.type === 'dragleave') {
      setDragActive(false);
    }
  };

  // Handle drop event
  const handleDrop = (e) => {
    e.preventDefault();
    e.stopPropagation();
    setDragActive(false);

    if (e.dataTransfer.files && e.dataTransfer.files.length > 0) {
      addSelectedFiles(Array.from(e.dataTransfer.files));
      logAnalyticsEvent('connect_select_files', { method: 'drag_drop' }).catch(() => {});
    }
  };

  // Handle typing "select" to open file browser
  const handleKeyDown = (e) => {
    if (e.key === 'Enter') {
      const input = e.target.value.trim().toLowerCase();
      if (input === 'select') {
        handleFileSelect();
        e.target.value = '';
      } else if (input) {
        addLog(`Command not recognized: ${input}`, 'error');
        e.target.value = '';
      }
    }
  };

  const tabs = [
    { name: title, icon: "⌘" },
    { name: isServer ? "Server" : "Client", icon: isServer ? "🖧" : "🖥" }
  ];

  return (
    <div
      className={`rounded-lg overflow-hidden border shadow-lg bg-terminal-dark min-h-[350px] sm:min-h-[400px] flex flex-col ${
        dragActive ? 'border-primary-400 shadow-glow' : 'border-terminal-border'
      }`}
      onDragEnter={handleDrag}
    >
      {/* Terminal Header */}
      <div className="bg-terminal-header flex items-center justify-between px-2 sm:px-4 py-2 border-b border-terminal-border">
        {/* Window Controls */}
        <div className="flex space-x-1.5 sm:space-x-2">
          <div className="w-2.5 h-2.5 sm:w-3 sm:h-3 rounded-full bg-red-500"></div>
          <div className="w-2.5 h-2.5 sm:w-3 sm:h-3 rounded-full bg-yellow-500"></div>
          <div className="w-2.5 h-2.5 sm:w-3 sm:h-3 rounded-full bg-green-500"></div>
        </div>

        {/* Tabs */}
        <div className="flex space-x-1">
          {tabs.map((tab, index) => (
            <div
              key={index}
              className="px-2 sm:px-3 py-1 text-xs rounded-t-md bg-terminal-dark text-terminal-light"
            >
              <span className="mr-1">{tab.icon}</span>
              <span className="hidden sm:inline">{tab.name}</span>
            </div>
          ))}
        </div>

        {/* Connection Status */}
        <div className="flex items-center space-x-2">
          <div className={`w-2 h-2 rounded-full ${isConnected ? 'bg-green-500' : 'bg-red-500'}`}></div>
          <span className="text-xs text-gray-400">{isConnected ? 'Connected' : 'Disconnected'}</span>
        </div>
      </div>

      {/* Terminal Content */}
      <div
        ref={terminalRef}
        className="p-3 sm:p-4 font-mono text-xs sm:text-sm overflow-y-auto flex-grow text-terminal-light"
        onDragOver={handleDrag}
        onDragLeave={handleDrag}
        onDrop={handleDrop}
      >
        {/* Logs */}
        <div className="mb-4">
          {logs.map((log, index) => (
            <div key={index} className={`mb-1 ${
              log.type === 'error' ? 'text-red-400' :
              log.type === 'success' ? 'text-green-400' :
              log.type === 'warning' ? 'text-yellow-400' :
              log.type === 'system' ? 'text-primary-400' :
              'text-terminal-light'
            }`}>
              <span className="text-gray-500">[{log.timestamp.toLocaleTimeString()}]</span> {log.message}
            </div>
          ))}
        </div>

        {/* Files List */}
        {files.length > 0 && (
          <div className="mb-4 border border-gray-800 rounded-md p-2">
            <div className="text-primary-400 mb-2">Files:</div>
            {files.map((file) => (
              <div key={file.id} className="flex items-center justify-between mb-1 text-xs">
                <div className="flex-1 truncate">{file.name}</div>
                <div className="text-gray-500 mx-2">{formatFileSize(file.size)}</div>
                <div className={`
                  ${file.status === 'queued' ? 'text-yellow-400' :
                    file.status === 'sending' ? 'text-primary-400' :
                    'text-green-400'}
                `}>
                  {file.status === 'queued' ? 'Queued' :
                   file.status === 'sending' ? 'Sending...' :
                   'Sent'}
                </div>
              </div>
            ))}
          </div>
        )}

        {/* Drag & Drop Overlay */}
        {dragActive && (
          <div className="absolute inset-0 bg-terminal-dark bg-opacity-90 flex items-center justify-center">
            <div className="text-center p-6 border-2 border-dashed border-primary-400 rounded-lg">
              <div className="text-primary-400 text-xl mb-2">Drop files here</div>
              <div className="text-gray-400">Release to add files</div>
            </div>
          </div>
        )}

        {/* Input Area */}
        <div className="mt-2 flex items-center">
          <span className="text-secondary-400 mr-1">$</span>
          <input
            type="text"
            className="flex-1 bg-transparent border-none outline-none text-terminal-light"
            placeholder="Type 'select' to open file browser or drag files here"
            onKeyDown={handleKeyDown}
          />
        </div>
      </div>

      {/* Controls */}
      <div className="bg-terminal-header border-t border-terminal-border p-2 sm:p-3 flex flex-wrap gap-2">
        <button
          onClick={handleConnect}
          className={`btn btn-sm ${isConnected ? 'bg-red-600 hover:bg-red-700' : 'bg-green-600 hover:bg-green-700'} text-white`}
          data-id="connect-toggle"
        >
          {isConnected ? 'Disconnect' : isServer ? 'Start Server' : 'Connect'}
        </button>

        <button
          onClick={handleFileSelect}
          className="btn btn-sm bg-primary-600 hover:bg-primary-700 text-white"
          disabled={!isConnected}
          data-id="select-files"
        >
          Select Files
        </button>

        <input
          type="file"
          ref={fileInputRef}
          className="hidden"
          onChange={handleFileInputChange}
          multiple
        />

        <div className="flex-1"></div>

        <div className="text-xs text-gray-400 flex items-center">
          {isServer ?
            `Server: Port ${port} • Save path: ${savePath}` :
            `Client: ${peerIp || 'No peer'} • Port: ${port}`
          }
        </div>
      </div>
    </div>
  );
}
