import {gql, useQuery} from "@apollo/client";
import Note from "../models/Note";

const DisplayNotes = () => {
    const GET_DATA = gql`
        query {
            getAllNotes {
                ID
                Content
            }
        }`;

  const { loading, error, data } = useQuery(GET_DATA);
  if (loading) return <p>Loading...</p>;
  if (error) return <p>Error :(</p>;

  return data.getAllNotes.slice(0, 5).map(({ ID, Content }: any) => (
    <div key={ID}>
      <Note
        ID={ID}
        Content={Content}
      />
    </div>
  ));
}

export default DisplayNotes;