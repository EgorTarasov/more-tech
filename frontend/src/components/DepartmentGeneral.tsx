import { Button, Col, Rate, Row, message } from 'antd';
import { IDepartment } from '../api/models';
import distanceConverter from '../utils/distanceConverter';
import { FileAddOutlined } from '@ant-design/icons';
import { useStores } from '../hooks/useStores';

type Props = {
    department: IDepartment;
};

const DepartmentGeneral = ({ department }: Props) => {
    const { rootStore } = useStores();
    const [messageApi, contextHolder] = message.useMessage();

    return (
        <>
            {contextHolder}
            <div
                onClick={() => {
                    rootStore.setSelectedDepartment(department);
                }}
                className='department-general'
            >
                <Row>
                    <div className='department-general__name department-general__item'>
                        {department.shortName}
                    </div>
                </Row>

                <Row>
                    <div className='department-general__wait-time department-general__item'>
                        Ожидание - 5 минут
                    </div>
                </Row>

                <Row wrap={false} justify={'space-between'}>
                    <Col span={18}>
                        <div className='department-general__address department-general__item'>
                            {department.address}
                        </div>
                    </Col>

                    <Col>
                        <div className='department-general__distance department-general__item'>
                            {distanceConverter(department.distance)}
                        </div>
                    </Col>
                </Row>

                <Row justify={'space-between'} align={'middle'} style={{ marginTop: 7 }}>
                    <Col>
                        <Rate
                            onChange={(value: number) =>
                                rootStore
                                    .postDepartmentRating(value, department._id)
                                    .then(() => {
                                        messageApi.success('Отзыв добавлен');
                                    })
                                    .catch(() => {
                                        messageApi.error('Ошибка добавления отзыва');
                                    })
                            }
                            allowHalf
                            defaultValue={department.rating}
                        />
                    </Col>

                    <Col>
                        В избранное
                        <Button
                            style={{
                                marginLeft: 10,
                                backgroundColor: department.favourite ? '#f5f5f5' : '#ebedf5',
                                color: '#6b7683',
                                border: 'none',
                            }}
                            type='default'
                            shape='circle'
                            onClick={() => {
                                rootStore
                                    .setAsFavorite(department._id)
                                    .then(() => {
                                        messageApi.success('Отделение добавлено в избранные');
                                    })
                                    .catch(() => {
                                        messageApi.error('Ошибка добавления отделения в избранное');
                                    });
                            }}
                            icon={<FileAddOutlined />}
                        />
                    </Col>
                </Row>
            </div>
        </>
    );
};

export default DepartmentGeneral;
