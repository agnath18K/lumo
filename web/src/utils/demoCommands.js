// Shared demo commands for use across components

// Hero section demo commands
export const heroCommands = [
  {
    input: 'lumo "find large files"',
    output: `find /path -type f -size +100M

Finds files larger than 100MB.`
  },
  {
    input: 'lumo auto:backup documents',
    output: `tar -czf backup.tar.gz ~/Documents

Created backup: documents_backup.tar.gz (156MB)`
  }
];

// Main demo section commands
export const mainDemoCommands = [
  {
    input: 'lumo "search for text in files"',
    output: `grep -r "pattern" /path

Search recursively for text in files.`
  },
  {
    input: 'lumo agent:find large files',
    output: `find . -type f -size +10M | sort -k5 -h

Found: video.mp4 (156M), image.png (42M), report.pdf (28M)`
  },
  {
    input: 'lumo health:system',
    output: `CPU: 12% | RAM: 5.2/16GB | Disk: 68% free

System health: Optimal`
  }
];

// Demo features for the Demo component
export const demoFeatures = [
  {
    title: "Natural Language Queries",
    description: "Ask questions in plain English and get the exact commands you need.",
    icon: 'ChatBubbleLeftRightIcon'
  },
  {
    title: "Terminal Agent Mode",
    description: "Let Lumo CLI execute sequences of terminal commands to complete complex tasks.",
    icon: 'CpuChipIcon'
  },
  {
    title: "System Health Monitoring",
    description: "Check system health and get detailed reports with simple commands.",
    icon: 'ChartBarIcon'
  }
];
