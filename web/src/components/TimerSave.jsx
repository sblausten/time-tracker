import React, {useState} from 'react';
import PropTypes from 'prop-types';

const TimerSave = (props) => {
    const [sessionName, setSessionName] = useState(undefined)

    const useHandleChange = (event) => {
        setSessionName(event.target.value);
    };

    const handleSubmit = (event) => {
        props.useSave(sessionName);
        event.preventDefault();
    };

    return (
        <div className="timer__save">
            <form onSubmit={handleSubmit}>
                <label htmlFor="session-name">Session Name</label>
                <input type="text" id="session-name" name="session-name" size="10" onChange={useHandleChange}/>
                <input name="save-session" type="submit" className="primary-button" value="Save" />
            </form>
        </div>
    )
};


TimerSave.propTypes = {
    useSave: PropTypes.func,
};

export default TimerSave;
