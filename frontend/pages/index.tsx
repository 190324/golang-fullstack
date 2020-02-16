
import { useQuery } from '@apollo/react-hooks'
// import { NetworkStatus } from 'apollo-client'
import gql from 'graphql-tag'

const TODOS = gql`
  query todos {
    todos {
        id
        text
    }
  }
`

const Page = () => {
    const { loading, error, data, fetchMore, networkStatus } = useQuery(
        TODOS,
        {
          notifyOnNetworkStatusChange: true,
        }
      )

    if (error) return <div>Error</div>
    if (loading) return <div>Loading</div>

    const { todos } = data

    console.log(todos);

    return (
        <section>
            <ul>
            {todos.map((todo, index) => (
                <li key={todo.id}>
                <div>
                    <span>{index + 1}. </span>
                    <span>{todo.text}</span>
                </div>
                </li>
            ))}
            </ul>
        </section>
    )
}

export default Page