import React from "react";
import PropTypes from "prop-types";

const UserId = (props) => {
    const {userId, setUserId} = props;

    const useHandleChange = (event) => {
        setUserId(event.target.value);
    };

    return (
        <div className="user">
            <div className="user__info">Your user id is: {userId}</div>
            <form id="user__update">
                <label htmlFor="user-id">Enter previous user id: </label>
                <input type="text" id="user__id" name="user-id" size="10" onChange={useHandleChange}/>
            </form>
        </div>
    )
};

UserId.propTypes = {
    userId: PropTypes.string,
    setUserId: PropTypes.func
};

export default UserId;