import axios from "axios";
import { useState } from "react";
import { useNavigate } from "react-router-dom"

export const Upload = () => {
  const navigate = useNavigate();
  const [dragging, setDragging] = useState(false);
  const [files, setFiles] = useState([]);
  const [url, setUrl] = useState('')

  const handleDragEnter = (e) => {
    e.preventDefault();
    setDragging(true);
  };

  const handleDragOver = (e) => {
    e.preventDefault();
    setDragging(true);
  };

  const handleDragLeave = () => {
    setDragging(false);
  };

  function addFiles(sFiles) {

    const validImageTypes = ['image/gif', 'image/jpeg', 'image/png'];

    if (sFiles.length < 1) {
      return
    }

    for (let i = 0; i < sFiles.length; i++) {
      if (validImageTypes.includes(sFiles[i]['type'])) {
        setFiles([...files, sFiles[i]])
      }
    }
  }

  function addFile(file) {

    const validImageTypes = ['image/gif', 'image/jpeg', 'image/png'];

    if (validImageTypes.includes(file['type'])) {
      setFiles([...files, file])
      return
    }

  }

  const handleDrop = (e) => {
    e.preventDefault();
    setDragging(false);
    addFiles(e.dataTransfer.files)
  };

  const handleBrowse = (e) => {
    let selectedFiles = e.target.files

    if (selectedFiles.length < 1) {
      return
    }

    addFile(selectedFiles[0])

  }

  // const handleAdd = async () => {
  //   try {
  //     const response = await axios.get(url, { responseType: 'stream' });
  //     addFile(response["data"])

  //   } catch (error) {
  //     console.error('Error downloading image:', error);
  //   }
  // }

  const handleSave = async () => {
    //TODO: limit upload to one or make enable bulk upload
    const url = "http://localhost:8080/upload/file"
    const form = new FormData();
    form.append('file', files[0]);
    await axios.post(url, form).then(
      _ => {
        console.log("");
        navigate('/collection')
      }
    ).catch(
      (error) => {
        console.log(error)
      }
    )
  }

  return (

    <div>
      <div className="flex flex-row m-6 p-1">
        <h1 className="text-2xl text-indigo-900 font-bold">Upload</h1>
        <div className="ml-6">
          <button onClick={handleSave} className="bg-indigo-400 rounded p-1">Save</button>
        </div>
      </div>
      <div className="max-w-xl m-5 mb-1"
        onDragEnter={handleDragEnter}
        onDragOver={handleDragOver}
        onDragLeave={handleDragLeave}
        onDrop={handleDrop}
      >
        <label
          className="flex justify-center w-full h-52 px-4 transition bg-white border-2 border-gray-300 border-dashed rounded-md appearance-none cursor-pointer hover:border-gray-400 focus:outline-none">
          <span className="flex items-center space-x-2">
            <svg xmlns="http://www.w3.org/2000/svg" className="w-6 h-6 text-gray-600" fill="none" viewBox="0 0 24 24"
              stroke="currentColor" stroke-width="2">
              <path stroke-linecap="round" stroke-linejoin="round"
                d="M7 16a4 4 0 01-.88-7.903A5 5 0 1115.9 6L16 6a5 5 0 011 9.9M15 13l-3-3m0 0l-3 3m3-3v12" />
            </svg>
            <span className="font-medium text-gray-600">
              Drop files to Attach, or{" "}
              <span className="text-blue-600 underline">browse</span>
            </span>
          </span>
          <input onChange={handleBrowse} type="file" name="file_upload" className="hidden" />
        </label>
      </div>
      {/* <div className="m-5 mt-1 mb-1"> */}
      {/*   <input value={url} */}
      {/*     type="text" */}
      {/*     className="bg-indigo-100 border-black center text-center text-indigo-900 p-1 pl-14 pr-14 m-1" */}
      {/*     placeholder="Link from the web" */}
      {/*     onChange={(e) => setUrl(e.target.value)} /> */}
      {/*   <button onClick={handleAdd} className="bg-indigo-400 rounded p-1">Add</button> */}

      {/* </div> */}


      <div className="flex flex-row m-5 mt-1">
        {files.map((file, key) => {
          return (
            <div key={key} className="overflow-hidden relative">
              <i onClick={() => { removeImage(file.name) }} className="mdi mdi-close absolute right-1 hover:text-white cursor-pointer"></i>
              <img className="h-20 w-20 rounded-md" src={URL.createObjectURL(file)} />
            </div>
          )
        })}

      </div>


    </div>
  );
};

