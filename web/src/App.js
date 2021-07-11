import './styles/App.css';
import React, {useState} from 'react';
import Timer from "./components/Timer";
import { v4 as uuidv4 } from 'uuid';

const App = () => {
    const [userId, setUserId] = useState(uuidv4());

    return (
        <div className="app">
            <div className="userInfo">Your session id is: {userId}</div>
            <Timer userId={userId}/>
        </div>
    );
};

export default App;
