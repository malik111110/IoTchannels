
import React, { useState, useEffect } from 'react';
import {
    Table,
    TableBody,
    TableCell,
    TableContainer,
    TableHead,
    TableRow,
    Paper,
    Typography,
    Card,
    CardContent,
} from '@mui/material';
import { styled } from '@mui/system';

const StyledTableRow = styled(TableRow)(({ theme }) => ({
    '&.fade-in': {
        animation: 'fadeIn 1s',
    },
    '@keyframes fadeIn': {
        from: {
            opacity: 0,
        },
        to: {
            opacity: 1,
        },
    },
}));

const Events = () => {
    const [events, setEvents] = useState([]);

    useEffect(() => {
        const socket = new WebSocket("ws://localhost:8080/ws");

        socket.onmessage = (event) => {
            const newEvent = JSON.parse(event.data);
            setEvents(prevEvents => [newEvent, ...prevEvents]);
        };

        return () => {
            socket.close();
        };
    }, []);

    return (
        <Card>
            <CardContent>
                <Typography variant="h5" component="div" sx={{ mb: 2 }}>
                    Events
                </Typography>
                <TableContainer component={Paper}>
                    <Table sx={{ minWidth: 650 }} aria-label="simple table">
                        <TableHead>
                            <TableRow>
                                <TableCell>Device ID</TableCell>
                                <TableCell>Type</TableCell>
                                <TableCell>Value</TableCell>
                                <TableCell>Time</TableCell>
                            </TableRow>
                        </TableHead>
                        <TableBody>
                            {events.map((event, index) => (
                                <StyledTableRow
                                    key={index}
                                    className={index === 0 ? 'fade-in' : ''}
                                >
                                    <TableCell component="th" scope="row">
                                        {event.DeviceID}
                                    </TableCell>
                                    <TableCell>{event.Type}</TableCell>
                                    <TableCell>{event.Value}</TableCell>
                                    <TableCell>{event.Time}</TableCell>
                                </StyledTableRow>
                            ))}
                        </TableBody>
                    </Table>
                </TableContainer>
            </CardContent>
        </Card>
    );
};

export default Events;
