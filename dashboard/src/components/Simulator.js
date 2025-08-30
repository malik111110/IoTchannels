
import React, { useState, useEffect } from 'react';
import {
    Button,
    Card,
    CardContent,
    Typography,
    Box,
    CircularProgress,
} from '@mui/material';
import { styled } from '@mui/system';

const PulsatingBox = styled(Box)(({ theme }) => ({
    width: 20,
    height: 20,
    borderRadius: '50%',
    backgroundColor: 'green',
    display: 'inline-block',
    marginLeft: theme.spacing(1),
    animation: 'pulse 2s infinite',
    '@keyframes pulse': {
        '0%': {
            boxShadow: '0 0 0 0 rgba(0, 255, 0, 0.7)',
        },
        '70%': {
            boxShadow: '0 0 0 10px rgba(0, 255, 0, 0)',
        },
        '100%': {
            boxShadow: '0 0 0 0 rgba(0, 255, 0, 0)',
        },
    },
}));

const Simulator = () => {
    const [running, setRunning] = useState(false);
    const [loading, setLoading] = useState(true);

    useEffect(() => {
        fetch('/api/status')
            .then(res => res.json())
            .then(data => {
                setRunning(data.running);
                setLoading(false);
            });
    }, []);

    const start = () => {
        setLoading(true);
        fetch('/api/start', { method: 'POST' })
            .then(() => {
                setRunning(true);
                setLoading(false);
            });
    };

    const stop = () => {
        setLoading(true);
        fetch('/api/stop', { method: 'POST' })
            .then(() => {
                setRunning(false);
                setLoading(false);
            });
    };

    return (
        <Card sx={{ mb: 4 }}>
            <CardContent>
                <Typography variant="h5" component="div">
                    Simulator Control
                </Typography>
                <Box sx={{ display: 'flex', alignItems: 'center', mt: 2 }}>
                    <Typography variant="body1">
                        Status: {running ? 'Running' : 'Stopped'}
                    </Typography>
                    {running && <PulsatingBox />}
                </Box>
                <Box sx={{ mt: 2 }}>
                    <Button variant="contained" onClick={start} disabled={running || loading} sx={{ mr: 2 }}>
                        {loading && !running ? <CircularProgress size={24} /> : 'Start'}
                    </Button>
                    <Button variant="contained" color="error" onClick={stop} disabled={!running || loading}>
                        {loading && running ? <CircularProgress size={24} /> : 'Stop'}
                    </Button>
                </Box>
            </CardContent>
        </Card>
    );
};

export default Simulator;
