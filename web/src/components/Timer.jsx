import React, {useState} from 'react';
import TimerControl from "./TimerControl";
import {client, handleErrorOnSave} from "../http/client";
import PropTypes from 'prop-types';
import TimerSave from "./TimerSave";

export const STOPPED = "Start";
export const STARTED = "Stop";

const Timer = (props) => {
    const { userId, isSaved, setIsSaved} = props
    const [timerState, setTimerState] = useState(STOPPED);
    const [startTime, setStartTime] = useState(undefined);
    const [stopTime, setStopTime] = useState(undefined);
    const [error, setError] = useState(null);
    const [isLoading, setIsLoading] = useState(false);

    const handleError = (error) => {
        setIsLoading(false);
        console.log("handleError: ", error);
        setError(error);
    };

    const useSave = (sessionName) => {
        function reset() {
            setIsSaved(false);
            setError(undefined);
            setIsLoading(true);
        }

        if (timerState === STOPPED) {
            reset();
            client.saveSession(userId, startTime, stopTime, sessionName)
                .then(handleErrorOnSave)
                .then(
                    result => {
                        setIsLoading(false);
                        console.log("Saved session: ", result);
                        setIsSaved(true);
                    },
                    networkError => {
                        setIsLoading(false);
                        console.log("NetworkError");
                        handleError(networkError);
                    }
                )
                .catch(error => {
                    console.log("Caught Error: ", error);
                    handleError(error);
                });

        }
    };


    return (
        <div className="timer">
            <TimerControl timerState={timerState} setIsSaved={setIsSaved} setError={setError}
                              setStartTime={setStartTime} setStopTime={setStopTime} setTimerState={setTimerState}/>
            <div className="timer__timings">
                {startTime && <div>Started at: {startTime.toString()}</div>}
                {stopTime && <div>Stopped at: {stopTime.toString()}</div>}
            </div>

            <TimerSave useSave={useSave}/>
            <div className="timer__status">
                {isLoading && "Loading..."}
                {isSaved && "Saved Session!"}
                {error && error.message}
            </div>
        </div>
    )
};

Timer.propTypes = {
    userId: PropTypes.string,
    isSaved: PropTypes.bool,
    setSessionSaved: PropTypes.func
};

export default Timer