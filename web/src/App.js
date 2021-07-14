import './styles/App.css';
import React, {useState} from 'react';
import Timer from "./components/Timer";
import SessionsList from "./components/SessionsList";
import { v4 as uuidv4 } from 'uuid';
import UserId from "./components/UserId";

const App = () => {
    const [userId, setUserId] = useState(uuidv4());
    const [isSaved, setIsSaved] = useState(false);

    return (
        <div className="app">
            <UserId setUserId={setUserId} userId={userId} />
            <Timer userId={userId} isSaved={isSaved} setIsSaved={setIsSaved}/>
            <SessionsList userId={userId} isSaved={isSaved} />
        </div>
    );
};

export default App;
