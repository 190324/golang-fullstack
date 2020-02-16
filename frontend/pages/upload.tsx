import gql from 'graphql-tag'
import { useMutation } from '@apollo/react-hooks'

export const SINGLE_UPLOAD = gql`
  mutation singleUpload($file: Upload!) {
    singleUpload(file: $file) {
      name
      content
    }
  }
`;

const Page = () => {

    const [uploadFile, {data}] = useMutation(SINGLE_UPLOAD);

    const onUpload = ({ target: { validity, files: [file] } }) => {
        validity.valid && uploadFile({variables: {file}});

        console.log(file);
    }

  return (   
      <div>
            <input
                type="file"
                accept="image/*"
                onChange={(e:any)=> {
                    onUpload(e);
                }}
            />

            {
                data && <div>
                    回應內容:<br/>
                    {data.singleUpload.name}<br/>
                    {data.singleUpload.content}
                </div>
            }
        </div>
  );
};

export default Page;