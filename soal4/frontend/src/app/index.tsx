import { Table, Input, Button, Image } from "antd";
import { useState } from "react";
import type { ColumnsType } from "antd/es/table";

interface User {
  key: string;
  name: string;
  age: number;
  location: string;
  email: string;
  phone: string;
  cell: string;
  picture: string;
}

const dummyData: User[] = [
  {
    key: "1",
    name: "Rudi Gunawan",
    age: 24,
    location: "Alamat Detail",
    email: "rudi@gmail.com",
    phone: "08227218291",
    cell: "08227218291",
    picture: "/profile1.jpg",
  },
  {
    key: "2",
    name: "Imam Ramadan",
    age: 27,
    location: "Alamat Detail",
    email: "imam@gmail.com",
    phone: "0824244424",
    cell: "0824244424",
    picture: "/profile2.jpg",
  },
];

export default function Home() {
  const [searchText, setSearchText] = useState("");
  const [filteredData, setFilteredData] = useState(dummyData);

  const handleSearch = (value: string) => {
    setSearchText(value);
    setFilteredData(
      dummyData.filter((user) =>
        user.name.toLowerCase().includes(value.toLowerCase())
      )
    );
  };

  const columns: ColumnsType<User> = [
    { title: "Nama", dataIndex: "name", key: "name" },
    { title: "Umur", dataIndex: "age", key: "age" },
    { title: "Alamat", dataIndex: "location", key: "location" },
    { title: "Email", dataIndex: "email", key: "email" },
    { title: "No. Telepon 1", dataIndex: "phone", key: "phone" },
    { title: "No. Telepon 2", dataIndex: "cell", key: "cell" },
    {
      title: "Gambar",
      dataIndex: "picture",
      key: "picture",
      render: (url) => <Image src={url} width={50} height={50} />,
    },
  ];

  return (
    <div style={{ padding: 20 }}>
      <h2>List</h2>
      <Input.Search
        placeholder="Search"
        value={searchText}
        onChange={(e) => handleSearch(e.target.value)}
        style={{ width: 300, marginBottom: 20 }}
      />
      <Table columns={columns} dataSource={filteredData} />
    </div>
  );
}
