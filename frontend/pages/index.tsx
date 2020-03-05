import { useQuery } from "@apollo/react-hooks";
// import { NetworkStatus } from 'apollo-client'
import gql from "graphql-tag";

const MEMBERS = gql`
  query members {
    members {
      data {
        sn
        first_name
      }
    }
  }
`;

const Page = () => {
  const { loading, error, data, fetchMore, networkStatus } = useQuery(MEMBERS, {
    // notifyOnNetworkStatusChange: true
  });

  if (error) return <div>Error</div>;
  if (loading) return <div>Loading</div>;

  const { members } = data;

  console.log(members);

  return (
    <section>
      <ul>
        {/* {members.map((todo, index) => (
          <li key={todo.id}>
            <div>
              <span>{index + 1}. </span>
              <span>{todo.text}</span>
            </div>
          </li>
        ))} */}
      </ul>
    </section>
  );
};

export default Page;
