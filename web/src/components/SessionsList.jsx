import React, {useEffect, useState} from 'react';
import {client, handleErrorOnSave} from "../http/client";
import moment from 'moment';

const SessionsList = (props) => {
    const {userId, isSaved} = props;
    const [sessions, setSessions] = useState([]);

    useEffect(() => {
        client.getSessions(userId)
            .then(res => res.json())
            .then(
                res => {
                    if(res.sessions) {
                        setSessions(res.sessions)
                    }
                },
                err => {
                    console.log(err);
                }
            )
    }, [isSaved]);

    const renderSessionRow = (session) => {
        const start = moment(session.start).format("DD/MM/YYYY h:mm:ss a");
        const end = moment(session.end).format("DD/MM/YYYY h:mm:ss a");
        const duration = moment.utc(moment.duration(session.duration).as('milliseconds')).format('HH[hr,] mm[m,] ss[s,] SSS[ms]');

        return (
            <tr className="session-row">
                <td>{start}</td>
                <td>{end}</td>
                <td>{session.name}</td>
                <td>{duration}</td>
            </tr>
        )
    };

    return (
        <div className="sessions-list">
            <h4>Your previous saved sessions:</h4>
            <table>
                <tr>
                    <th>Start</th>
                    <th>End</th>
                    <th>Name</th>
                    <th>Duration</th>
                </tr>
                {sessions && sessions.map(s => renderSessionRow(s))}
            </table>
        </div>
    )

};

export default SessionsList;
