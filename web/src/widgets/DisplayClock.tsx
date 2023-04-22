import { useState } from "react";

const DisplayClock = () => {
  const [time, setTime] = useState(new Date().toLocaleTimeString());

  const UpdateTime = () => {
    setTime(new Date().toLocaleTimeString());
  }
  setInterval(UpdateTime);
  return <h1>{time}</h1>
}

export default DisplayClock;