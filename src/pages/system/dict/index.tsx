import type {ActionType, ProColumns} from '@ant-design/pro-components';
import { ProCard, ProTable } from '@ant-design/pro-components';
import type { BadgeProps } from 'antd';
import { Badge, Button } from 'antd';
import React, {useEffect, useRef, useState} from 'react';
import UpdateForm from "@/pages/system/dict/list/components/UpdateForm";
import {getDicthtList, getDictList} from "@/services/ant-design-pro/dict";
import {getUser} from "@/services/ant-design-pro/user";
import {history} from "@@/core/history";

type DicthtListProps = {
  id: number;
};

const DicthtList: React.FC<DicthtListProps> = (props) => {
  const { id  } = props;

  const [dicthtId, setDicthtId] = useState<number>(0);

  setDicthtId(id)

  const [searchDicthtName, setSearchDicthtName] = useState<string>('');

  const columns: ProColumns<API.Dictht>[] = [

    {
      title: "标题",
      dataIndex: 'title',
      sorter: true,
      hideInForm: true,
    },
    {
      title: "key",
      dataIndex: 'key',
      sorter: true,
      hideInForm: true,
    },
    {
      title: "value",
      dataIndex: 'value',
      sorter: true,
      hideInForm: true,
    },
    {
      title: '状态',
      dataIndex: 'status',
      hideInForm: true,
      valueEnum: {
        0:{
          text: '禁用',
          status: 'Error',
        },
        1: {
          text: '正常',
          status: 'Success',
        },

      },
    },
    {
      title: '操作',
      key: 'option',
      width: 80,
      valueType: 'option',
      render: () => [<a key="a">预警</a>],
    },
  ];

  // const fetchDicthtList = async () => {
  //   try {
  //     console.log(id)
  //     const msg = await getDicthtList({dictId:id});
  //     setTableListDataSource(msg.data);
  //   } catch (_error) {
  //
  //   }
  // }
  //
  // useEffect(() => {
  //  fetchDicthtList();
  // }, [id]);

  return (
    <ProTable
      columns={columns}
      rowKey="id"
      search={false}
      params={{dictId: dicthtId, name: searchDicthtName}}
      request={getDicthtList}
      toolbar={{
        search: {
          onSearch: (value) => {
            setSearchDicthtName(value);
          },
        },
        actions: [
          <Button key="list" type="primary">
            新建
          </Button>,
        ],
      }}
    />
  );
};


type DictListProps = {
  id?: number;
  onChange: (id: number) => void;
};


const DictList: React.FC<DictListProps> = (props) => {
  const { onChange } = props;

  const [searchDictName, setSearchDictName] = useState<string>('');


  const actionRef = useRef<ActionType | null>(null);

  const [showDetail, setShowDetail] = useState<boolean>(false);
  const [currentRow, setCurrentRow] = useState<API.Dict>();
  const [selectedRowsState, setSelectedRows] = useState<API.Dict[]>([]);


  const columns: ProColumns<API.Dict>[] = [
    {
      title: '名称',
      dataIndex: 'name',
      render: (dom, entity) => {
        return (
          <a
            onClick={() => {
              setCurrentRow(entity);
              setShowDetail(true);
            }}
          >
            {dom}
          </a>
        );
      },
    },
    {
      title: "标题",
      dataIndex: 'title',
    },
    {
      title: "概略",
      dataIndex: 'desc',
      valueType: 'textarea',
    },
    {
      title: '状态',
      dataIndex: 'status',
      valueEnum: {
        0:{
          text: '禁用',
          status: 'Error',
        },
        1: {
          text: '正常',
          status: 'Success',
        },

      },
    },
    {
      title: '操作',
      dataIndex: 'option',
      valueType: 'option',
      render: (_, record) => [
        <UpdateForm
          trigger={
            <a>更新</a>
          }
          key="config"
          onOk={actionRef.current?.reload}
          values={record}
        />,
      ],
    },
  ];



  return (
    <ProTable<API.Dict>
      columns={columns}
      params={{name: searchDictName}}
      request={getDictList}
      rowKey="id"
      toolbar={{
        search: {
          onSearch: (value) => {
            setSearchDictName(value);
          },
        },

        actions: [
          <Button key="list" type="primary">
            新建
          </Button>,
        ],
      }}
      // options={false}
      // pagination={false}
      search={false}
      onRow={(record) => {
        return {
          onClick: () => {
            if (record.id) {
              console.log(record.id)
              onChange(record.id);
            }
          },
        };
      }}
    />
  );
};

const Pro: React.FC = () => {
  const [id, setId] = useState(0);
  return (
    <ProCard split="vertical">
      <ProCard colSpan="40%" ghost>
        <DictList  onChange={(cId) =>setId(cId) }  id={id} />
      </ProCard>
      <ProCard>
        <DicthtList  id={id}  />
      </ProCard>
    </ProCard>
  );
};

const Pages = () => {
  return (
    <>
      <Pro />
    </>
  );
};

export default Pages;
