import '../css/Reminder.scss';

interface ReminderProps {
  id: string;
  title: string;
  description: string;
  date: Date;
  isDone: boolean;
}

const Reminder = (props: ReminderProps) => {
  return (
    <>
      <tr className="reminder-item">
        <td className="reminder-item__title">{props.title}</td>
        <td className="reminder-item__description">{props.description}</td>
        <td>{props.date.toString()}</td>
      </tr>
    </>
  )
}

export default Reminder;