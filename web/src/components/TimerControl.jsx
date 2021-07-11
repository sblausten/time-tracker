import React from 'react';
import PropTypes from 'prop-types';
import {STARTED, STOPPED} from "./Timer";

const TimerControl = (props) => {
    const {timerState, setIsSaved, setError, setStartTime, setStopTime, setTimerState} = props;
    const classNames = "primary-button";

    const useStartStopTimer = () => {
        const now = new Date();

        if (timerState === STOPPED) {
            setIsSaved(false);
            setError(null);
            setStartTime(now);
            setStopTime(undefined);
            setTimerState(STARTED);
        } else if (timerState === STARTED) {
            setStopTime(now);
            setTimerState(STOPPED);
        }
    };

    return (
        <div className="timer__control">
            <button name={timerState} type="button" className={classNames} onClick={useStartStopTimer}>
                {timerState}
            </button>
        </div>
    );
};

TimerControl.propTypes = {
    timerState: PropTypes.string,
    setIsSaved: PropTypes.func,
    setError: PropTypes.func,
    setStartTime: PropTypes.func,
    setStopTime: PropTypes.func,
    setTimerState: PropTypes.func
};

export default TimerControl;
