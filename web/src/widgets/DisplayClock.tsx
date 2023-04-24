import { useState } from "react";
import '../css/Clock.scss';

const DisplayClock = () => {
  const [time, setTime] = useState(new Date().toLocaleTimeString());

  const UpdateTime = () => {
    setTime(new Date().toLocaleTimeString());
  }
  setInterval(UpdateTime);
  return (
    <div className={"clock"}>
      <h1>{time}</h1>
    </div>
    )
}

export default DisplayClock;