import {useEffect, useState} from "react";

import './App.scss';
import DisplayClock from "./widgets/DisplayClock";
import DisplayFlashcards from "./widgets/DisplayFlashcards";
import DisplayJoke from './widgets/DisplayJoke';
import DisplayNotes from "./widgets/DisplayNotes";
import DisplayReminders from "./widgets/DisplayReminders";
import DisplayTask from "./widgets/DisplayTask";
import DisplayNumbers from "./widgets/DisplayNumbers";
import DisplayPicture from "./widgets/DisplayPicture";
function App() {
  const [state, refreshState] = useState(false);

  useEffect(() => {
    setInterval(() => {
      refreshState(!state);
    }, 86400000);
  }, [state]);

  return (
    <div className="App">
      <div className="box">
        <DisplayClock />
      </div>
      <div className="box">
        <DisplayReminders />
      </div>
      <div className="box">
        <DisplayFlashcards />
      </div>
      <div className="box">
        <DisplayNotes />
      </div>
      <div className="box">
        <DisplayJoke />
      </div>
      <div className="box">
        <DisplayTask />
      </div>
      <div className="box">
        <DisplayNumbers />
      </div>
      <div className="box">
        <DisplayPicture />
      </div>
    </div>
  );
}

export default App;