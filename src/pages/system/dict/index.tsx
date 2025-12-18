import type {ActionType, ProColumns} from '@ant-design/pro-components';
import { ProCard, ProTable } from '@ant-design/pro-components';
import React, {useEffect, useRef, useState} from 'react';

import CreateForm from './components/CreateForm';
import UpdateForm from './components/UpdateForm';

import { Dictht,Dict } from  "@/pages/system/dict/service/data";
import {getDicthtList, getDictList}  from "@/pages/system/dict/service/service";

import CreateDicthtForm from "@/pages/system/dict/components/CreateDicthtForm";
import UpdateDicthtForm from "@/pages/system/dict/components/UpdateDicthtForm";

type DicthtListProps = {
  id: number;
};

const DicthtList: React.FC<DicthtListProps> = (props) => {
  const { id  } = props;

  const [dicthtId, setDicthtId] = useState<number>(0);

  const actionRef = useRef<ActionType | null>(null);

  const [searchDicthtName, setSearchDicthtName] = useState<string>('');

  const columns: ProColumns<Dictht>[] = [
    {
      title: "id",
      dataIndex: 'id',
      sorter: true,
      hideInForm: true,
    },
    {
      title: "标题",
      dataIndex: 'title',
      sorter: true,
      hideInForm: true,
    },

    {
      title: "有效值",
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
      dataIndex: 'option',
      valueType: 'option',
      render: (_, record) => [
        <UpdateDicthtForm
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
  useEffect(() => {
    setDicthtId(id)
  }, [id]);

  return (
    <ProTable
      columns={columns}
      rowKey="id"
      search={false}
      actionRef={actionRef}
      params={{dictId: dicthtId, key: searchDicthtName}}
      request={getDicthtList}
      toolbar={{
        search: {
          onSearch: (value) => {
            setSearchDicthtName(value);
          },
        },
        actions: [
          <CreateDicthtForm key="create" reload={actionRef.current?.reload} dictId={id} />,
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
  const [currentRow, setCurrentRow] = useState<Dict>();
  const [selectedRowsState, setSelectedRows] = useState<Dict[]>([]);


  const columns: ProColumns<Dict>[] = [
    {
      title: "id",
      dataIndex: 'id',
      sorter: true,
      hideInForm: true,
    },
    {
      title: '标识',
      dataIndex: 'code',
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
    <ProTable<Dict>
      columns={columns}
      params={{key: searchDictName}}
      request={getDictList}
      actionRef={actionRef}
      rowKey="id"
      toolbar={{
        search: {
          onSearch: (value) => {
            setSearchDictName(value);
          },
        },

        actions: [
          <CreateForm key="create" reload={actionRef.current?.reload} />,
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
        <DictList onChange={(cId) =>setId(cId) }  id={id} />
      </ProCard>
      <ProCard>
        <DicthtList id={id}  />
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
