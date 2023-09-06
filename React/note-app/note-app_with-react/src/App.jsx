import './App.css';
import Sidebar from './components/Sidebar';
import Main from './components/Main';
import { useEffect, useState } from 'react';
import uuid from 'react-uuid';

function App() {
  const [notes, setNotes] = useState(JSON.parse(localStorage.getItem('notes')) || []);
  const [sidebarActive, setSidebarActive] = useState(true);
  const [activeNote, setActiveNote] = useState(false);

  const onSidebarIsActive = () => {
    setSidebarActive(!sidebarActive);
  }

  const onAddNote = () => {
    const newNote = {
      id: uuid(),
      title: "新しいノート",
      content: "",
      modDate: Date.now()
    }
    setNotes([...notes, newNote]);
    console.log(notes);
  }

  const onDeleteNote = (id) => {
    console.log(id);
    const newNotes = notes.filter((note) => {
      return note.id !== id;
    });
    setNotes(newNotes);
  }

  const getActiveNote = (id) => {
    return notes.find((note) => note.id === id);
  }

  const onUpdatedNote = (updatedNote) => {
    const updatedNotesArray = notes.map((note) => {
      if(note.id === updatedNote.id) {
        return updatedNote;
      } else {
        return note;
      }
    })

    setNotes(updatedNotesArray);
  }

  useEffect(() => {
    localStorage.setItem('notes', JSON.stringify(notes));
  }, [notes]);

  useEffect(() => {
    setActiveNote(notes[0].id);
  }, []);

  return <div className="App">
    <Sidebar
      onAddNote={onAddNote}
      notes={notes}
      sidebarActive={sidebarActive}
      onSidebarIsActive={onSidebarIsActive}
      onDeleteNote={onDeleteNote}
      activeNote={activeNote}
      setActiveNote={setActiveNote}
    />
    <Main
      activeNote={getActiveNote(activeNote)}
      notes={notes}
      setNotes={setNotes}
      onUpdatedNote={onUpdatedNote}
    />
  </div>
}

export default App
