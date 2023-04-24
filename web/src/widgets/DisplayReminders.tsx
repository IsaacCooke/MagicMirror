import {gql, useQuery} from "@apollo/client";
import Reminder from "../models/Reminder";
import '../css/Reminder.scss';

const DisplayReminders = () => {
    const GET_DATA = gql`
        query {
            getAllReminders {
                ID
                IsDone
                Title
                Description
                DueDate
            }
        }`;

  const { loading, error, data } = useQuery(GET_DATA);
  if (loading) return <p>Loading...</p>;
  if (error) return <p>Error :(</p>;

  const ParseDate = (inputDate: Date): string => {
    const date = new Date(inputDate);

    const time = date.toLocaleTimeString([], {hour: '2-digit', minute:'2-digit'});
    const dateFormatted = date.toLocaleDateString([], {day: "2-digit", month: "2-digit", year: "numeric"});

    return `${time} ${dateFormatted}`;
  }

  const GetReminderRows = () => {
    return data.getAllReminders.slice(0, 5).map(({ ID, Title, IsDone, Description, DueDate }: any) => (
      <div key={ID}>
        <Reminder
          id={ID}
          title={Title}
          description={Description}
          isDone={IsDone}
          date={ParseDate(DueDate)}
        />
      </div>
    ));
  }

  return (
    <div className={"reminder-container"}>
      <h1 className={"reminder-title"}>Reminders</h1>
      <table className={"reminder-table"}>
        <tbody>
          <GetReminderRows />
        </tbody>
      </table>
    </div>
  )
}

export default DisplayReminders;