import {gql, useQuery} from "@apollo/client";
import Reminder from "../models/Reminder";

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

  return data.getAllReminders.slice(0, 5).map(({ ID, Title, IsDone, Description, DueDate }: any) => (
    <div key={ID}>
      <Reminder
        id={ID}
        title={Title}
        description={Description}
        isDone={IsDone}
        date={DueDate}
      />
    </div>
  ));
}

export default DisplayReminders;