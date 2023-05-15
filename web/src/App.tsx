import './App.scss'
import DisplayClock from "./widgets/DisplayClock";
import DisplayFlashcards from "./widgets/DisplayFlashcards";
import DisplayJoke from './widgets/DisplayJoke';
import DisplayNotes from "./widgets/DisplayNotes";
import DisplayReminders from "./widgets/DisplayReminders";
import DisplayTask from "./widgets/DisplayTask";

function App() {
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
    </div>
  );
}

export default App;