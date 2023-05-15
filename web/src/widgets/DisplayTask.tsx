import {useState, useEffect} from "react";
import '../css/Api.scss';

const DisplayTask = () => {
  const [state, refreshState] = useState(false);
  const [task, setTask] = useState("");

  useEffect(() => {
    setInterval(() => {
      refreshState(!state);
    }, 300000);

    fetch("https://www.boredapi.com/api/activity", {
      headers: {
        Accept: "application/json"
      }
    })
      .then(response => response.json())
      .then(data => setTask(data.activity));
  }, [state])

  return (
    <div className={"api-container"}>
      <h1 className={"api-text"}>{task}</h1>
    </div>
  )
}

export default DisplayTask;