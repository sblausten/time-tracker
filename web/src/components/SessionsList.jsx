import React, {useEffect, useState} from 'react';
import {client, handleErrors} from "../http/client";


const SessionsList = (props) => {
    const {userId, isSaved} = props;
    const [sessions, setSessions] = useState([]);

    useEffect(() => {
        client.getSessions(userId)
            .then(res => res.json())
            .then(
                res => {
                    setSessions(res.sessions)
                },
                err => {
                    console.log(err);
                }
            )
    }, [isSaved]);

    const renderSessionRow = (session) => {

        return (
            <div className="session-row">
                {session.start} - {session.end} || Duration: {session.duration / 1000 / 60} minutes{session.name && ", " + session.name}
            </div>
        )
    };

    return (
        <div className="sessions-list">
            <h4>Your previous saved sessions:</h4>
            <ul>
                {sessions.map(s => renderSessionRow(s))}
            </ul>
        </div>
    )

};

export default SessionsList;
