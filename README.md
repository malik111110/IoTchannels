# IoT Device Monitoring Dashboard

A real-time dashboard for monitoring IoT devices, built with React and WebSockets. This application provides a clean, responsive interface to visualize and interact with various IoT sensor data including temperature, motion detection, and door status.

## Features

- Real-time data visualization
- Support for multiple device types (temperature, motion, door sensors)
- Responsive design for desktop and mobile
- Historical data display
- Device status indicators
- WebSocket integration for live updates

## Prerequisites

- Node.js (v14 or higher)
- npm (v6 or higher) or Yarn
- Modern web browser (Chrome, Firefox, Safari, or Edge)

## Getting Started

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/chanIoT.git
   cd chanIoT
   ```

2. Install dependencies for both server and client:
   ```bash
   # Install server dependencies
   cd server
   npm install
   
   # Install client dependencies
   cd ../dashboard
   npm install
   ```

### Running the Application

1. Start the WebSocket server:
   ```bash
   cd server
   npm start
   ```

2. In a new terminal, start the React development server:
   ```bash
   cd dashboard
   npm start
   ```

3. Open [http://localhost:3000](http://localhost:3000) to view the dashboard in your browser.

## Project Structure

```
chanIoT/
├── dashboard/           # React frontend application
│   ├── public/         # Static files
│   └── src/            # React components and logic
├── server/             # WebSocket server
│   ├── src/            # Server source code
│   └── package.json    # Server dependencies
└── README.md           # This file
```

## Available Scripts

In the project directory, you can run:

### `npm start`
Runs the app in development mode.

### `npm test`
Launches the test runner in interactive watch mode.

### `npm run build`
Builds the app for production to the `build` folder.

## Environment Variables

Create a `.env` file in the root directory with the following variables:

```
REACT_APP_WS_URL=ws://localhost:8080
PORT=3000
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgements

- [React](https://reactjs.org/)
- [WebSocket API](https://developer.mozilla.org/en-US/docs/Web/API/WebSockets_API)
- [Chart.js](https://www.chartjs.org/) for data visualization
