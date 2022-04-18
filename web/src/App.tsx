import "./App.css";
import { gql, useQuery } from "@apollo/client";

type Gender = "F" | "M";
type Employee = {
  empNo: string;
  birthDate: string;
  firstName: string;
  lastName: string;
  gender: Gender;
  hireDate: string;
};
function App() {
  const { loading, error, data } = useQuery<{
    employees: Employee[];
  }>(gql`
    query {
      employees(limit: 10, offset: 0) {
        empNo
        firstName
        lastName
      }
    }
  `);
  if (loading) {
    return <p>Loading...</p>;
  }
  if (error) {
    return <p>Error :(</p>;
  }
  return (
    <div className="App">
      {data?.employees.map((employee) => (
        <div key={employee.empNo}>
          <h1>{employee.firstName}</h1>
          <p>{employee.lastName}</p>
        </div>
      ))}
    </div>
  );
}

export default App;
