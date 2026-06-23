import { useState, useEffect, useRef } from 'react';

export default function Terminal({
  commands = [],
  autoPlay = false,
  loop = false,
  typingSpeed = 20,  // Faster typing speed
  waitTime = 1500,   // Shorter wait time
  showControls = true,
  title = "Terminal"
}) {
  const [currentCommandIndex, setCurrentCommandIndex] = useState(0);
  const [displayedText, setDisplayedText] = useState('');
  const [isTyping, setIsTyping] = useState(false);
  const [showCursor, setShowCursor] = useState(true);
  const [history, setHistory] = useState([]);
  const [activeTab, setActiveTab] = useState(0);
  const [isInitialized, setIsInitialized] = useState(false);
  const [isCompleted, setIsCompleted] = useState(false);
  const [isLooping, setIsLooping] = useState(false);
  const terminalRef = useRef(null);
  const commandsRef = useRef(commands);

  // Update commands ref when commands prop changes
  useEffect(() => {
    commandsRef.current = commands;
  }, [commands]);

  // Initialize with a delay to ensure component is fully mounted
  useEffect(() => {
    if (autoPlay && !isInitialized) {
      const initTimer = setTimeout(() => {
        setIsInitialized(true);
      }, 500); // Reduced from 800ms
      return () => clearTimeout(initTimer);
    }
  }, [autoPlay, isInitialized]);

  // Cursor blink effect
  useEffect(() => {
    const cursorInterval = setInterval(() => {
      setShowCursor(prev => !prev);
    }, 500);

    return () => clearInterval(cursorInterval);
  }, []);

  // Main command execution loop
  useEffect(() => {
    if (autoPlay && commands.length > 0 && !isTyping && isInitialized && !isLooping) {
      const timer = setTimeout(() => {
        typeCommand(commands[currentCommandIndex]);
      }, waitTime);

      return () => clearTimeout(timer);
    }
  }, [currentCommandIndex, commands, autoPlay, isTyping, isInitialized, isLooping]);

  // Handle completion and looping
  useEffect(() => {
    if (isCompleted && loop && !isLooping) {
      setIsLooping(true);

      // Add a delay before resetting to show all commands
      const loopTimer = setTimeout(() => {
        setCurrentCommandIndex(0);
        setHistory([]);
        setIsCompleted(false);
        setIsLooping(false);
      }, 2000); // Reduced from 3000ms

      return () => clearTimeout(loopTimer);
    }
  }, [isCompleted, loop, isLooping]);

  // Scroll to bottom when history changes
  useEffect(() => {
    if (terminalRef.current) {
      terminalRef.current.scrollTop = terminalRef.current.scrollHeight;
    }
  }, [history]);

  const typeCommand = (command) => {
    setIsTyping(true);

    // Start with the first character already displayed
    setDisplayedText(command.input.charAt(0));

    // Add a small delay before continuing to type
    setTimeout(() => {
      let i = 1; // Start from the second character
      const typingInterval = setInterval(() => {
        setDisplayedText((prev) => prev + command.input.charAt(i));
        i++;

        if (i >= command.input.length) {
          clearInterval(typingInterval);

          // Add command to history after typing is complete
          setTimeout(() => {
            // Use functional updates to ensure state is properly updated
            setHistory(prev => [...prev, { type: 'input', text: command.input }]);

            // Add a small delay before showing output
            setTimeout(() => {
              setHistory(prev => [...prev, { type: 'output', text: command.output }]);
              setDisplayedText('');
              setIsTyping(false);

              // Check if this was the last command
              if (currentCommandIndex >= commandsRef.current.length - 1) {
                setIsCompleted(true);
              } else {
                // Move to next command
                setTimeout(() => {
                  setCurrentCommandIndex(prev => prev + 1);
                }, 300);
              }
            }, 300);
          }, waitTime);
        }
      }, typingSpeed);
    }, 300);
  };

  const handleReset = () => {
    setHistory([]);
    setCurrentCommandIndex(0);
    setIsTyping(false);
    setDisplayedText('');
    setIsCompleted(false);
    setIsLooping(false);
  };

  const tabs = [
    { name: title, icon: "⌘" },
    { name: "bash", icon: ">" },
    { name: "zsh", icon: "%" }
  ];

  return (
    <div className="rounded-lg overflow-hidden border border-terminal-border shadow-lg bg-terminal-dark min-h-[250px] sm:min-h-[300px] flex flex-col">
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
            <button
              key={index}
              onClick={() => setActiveTab(index)}
              className={`px-2 sm:px-3 py-1 text-xs rounded-t-md transition-colors ${
                activeTab === index
                  ? 'bg-terminal-dark text-terminal-light'
                  : 'bg-terminal-header text-gray-500 hover:text-gray-300'
              }`}
            >
              <span className="mr-1">{tab.icon}</span>
              <span className="hidden sm:inline">{tab.name}</span>
            </button>
          ))}
        </div>

        {/* Controls */}
        {showControls && (
          <div className="flex space-x-2">
            <button
              onClick={handleReset}
              className="text-gray-400 hover:text-gray-200 transition-colors p-1 min-w-[32px] min-h-[32px]"
              title="Clear terminal"
            >
              <svg xmlns="http://www.w3.org/2000/svg" className="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16" />
              </svg>
            </button>
          </div>
        )}
      </div>

      {/* Terminal Content */}
      <div
        ref={terminalRef}
        className="p-3 sm:p-4 font-mono text-xs sm:text-sm overflow-y-auto flex-grow text-terminal-light"
      >
        {history.map((item, index) => (
          <div key={index} className={`${item.type === 'input' ? 'text-primary-400' : 'text-terminal-light'} ${item.type === 'output' ? 'mt-1 mb-3 sm:mb-4 pl-2 sm:pl-4 border-l-2 border-gray-800' : ''}`}>
            {item.type === 'input' && <span className="text-secondary-400">$ </span>}
            {item.text}
          </div>
        ))}
        {isTyping && (
          <div className="text-primary-400">
            <span className="text-secondary-400">$ </span>
            {displayedText}
            {showCursor && <span className="inline-block w-1.5 sm:w-2 h-3 sm:h-4 bg-primary-400 ml-0.5 animate-terminal-cursor"></span>}
          </div>
        )}
        {!isTyping && !autoPlay && currentCommandIndex < commands.length && (
          <button
            onClick={() => typeCommand(commands[currentCommandIndex])}
            className="bg-primary-600 text-white px-3 py-2 rounded text-xs mt-2 hover:bg-primary-700 transition-colors focus:outline-none focus:ring-2 focus:ring-primary-500 focus:ring-offset-2 focus:ring-offset-terminal-dark min-h-[36px] min-w-[44px]"
          >
            Run next command
          </button>
        )}
      </div>
    </div>
  );
}
