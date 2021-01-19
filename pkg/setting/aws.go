package setting

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/aws/aws-sdk-go/service/s3"
	"my_gin/pkg/util"
	"os"
	"reflect"
	"strings"
)

type AwsS3Config struct {
	Region    string `json:"region"`
	Version   string `json:"version"`
	Endpoint  string `json:"endpoint"`
	AccessKey string `json:"access_key"`
	SecretKey string `json:"secret_key"`
	Bucket    string `json:"bucket"`
}

type AwsDynamoDbConfig struct {
	Region      string `json:"region"`
	Version     string `json:"version"`
	Endpoint    string `json:"endpoint"`
	AccessKey   string `json:"access_key"`
	SecretKey   string `json:"secret_key"`
	TablePrefix string `json:"table_prefix"`
}

type AwsCfg struct {
	S3Init bool
	S3Cfg  *AwsS3Config `json:"s3"`
	S3Client *s3.S3
	DynamoDbInit bool
	DynamoDbClient dynamodbiface.DynamoDBAPI
	DynamoDbCfg  *AwsDynamoDbConfig `json:"dynamodb"`
}

//dynamodb需要在命令行配置如下：
//E:\wwwroot\home_backend_go\backend\trunk>aws configure
//AWS Access Key ID [None]: AKIAOWWDRSMP3YGLMHSA
//AWS Secret Access Key [None]: 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
//Default region name [None]: cn-north-1
//Default output format [None]: json


func (this *AwsCfg)LoadAwsCfg(env string) error {
	file := util.NewFile()
	awsFile := fmt.Sprintf("conf/%s/aws.json", env)
	jsonStr, _ := file.GetContentString(awsFile)
	err := json.Unmarshal([]byte(jsonStr), this)
	if err != nil {
		fmt.Printf("Could Unmarshal %s: %s\n", jsonStr, err)
		return err
	}

	if err = this.InitAwsS3(); err != nil {
		fmt.Printf("InitAwsS3 error, err:%v!", err)
		return err
	}
	if err = this.InitDynamoDb(); err != nil {
		fmt.Printf("InitDynamoDb error, err:%v!", err)
		return err
	}
	return nil
}

func (this *AwsCfg)InitAwsS3() error{
	if this.S3Init {
		return nil
	}
	s3Session, err := session.NewSession(&aws.Config{
		Credentials: credentials.NewStaticCredentials(this.S3Cfg.AccessKey, this.S3Cfg.SecretKey, ""),
		Endpoint: aws.String(this.S3Cfg.Endpoint),
		Region: aws.String(this.S3Cfg.Region),
		DisableSSL: aws.Bool(true),
		S3ForcePathStyle: aws.Bool(false),
	})
	if err != nil {
		fmt.Printf("InitAwsS3 error, err:%v!", err)
		return err
	}
	s3Client := s3.New(s3Session)
	this.S3Client = s3Client
	return nil
}

func (this *AwsCfg)InitDynamoDb() error{
	if this.DynamoDbInit {
		return nil
	}
	awscfg := &aws.Config{}
	awscfg.WithRegion(this.DynamoDbCfg.Region)
	sess, err := session.NewSession(awscfg)
	if err != nil {
		panic(fmt.Errorf("failed to create session, %v", err))
	}
	dynamoDbClient := dynamodb.New(sess)
	this.DynamoDbClient = dynamoDbClient
	return nil
}

//获取S3 key里的内容
func (this *AwsCfg)S3GetObject(key string)(int64, []byte){
	params := &s3.GetObjectInput{
		Bucket: aws.String(this.S3Cfg.Bucket),
		Key: &key,
	}
	result,_ := this.S3Client.GetObject(params)
	buf := new(bytes.Buffer)
	length, _ := buf.ReadFrom(result.Body)
	content := buf.Bytes()
	return length, content
}

//更新S3的文件内容
func (this *AwsCfg)S3PutObject(key string, file *os.File)(string, error){
	params := &s3.PutObjectInput{
		Bucket: aws.String(this.S3Cfg.Bucket),
		Key: &key,
		Body: file,
	}
	context.Background()
	result,err := this.S3Client.PutObject(params)

	if err != nil {
		return "", err
	}
	return *result.ETag,nil
}


//根据主键查询数据【当主键没有排序键时使用】
//如果查询的表主键有排序件，推荐使用DynamoDbGetItemByKvs()
//@param ProjectionExpression 需要返回的字段“,”隔开， 如： key1,key2,key3
func (this *AwsCfg)DynamoDbGetItem(table, key, value, ProjectionExpression string)(*dynamodb.GetItemOutput, error){
	tableName := this.DynamoDbCfg.TablePrefix + table
	params := &dynamodb.GetItemInput{
		TableName: &tableName,
		Key: map[string]*dynamodb.AttributeValue{
			key : {
				S: aws.String(value),
			},
		},
	}
	if ProjectionExpression != "" {
		params.ProjectionExpression = aws.String(ProjectionExpression)
	}
	result, err := this.DynamoDbClient.GetItem(params)
	return result, err
}


// 根据【主键】查询数据【推荐使用】 aws后台视图模式中：表名处会显示索引列表（包括主键）
//如：[表] Gardenscapes_dev_Chat: chatId, msgId
//    [索引] FromIdIndex: fromId, type
//@param table 表名
//@param kvs 查询的条件数组， 如： array("主键" => '条件值1', '二级索引项目键' => "值")
//@param ProjectionExpression 需要返回的字段“,”隔开， 如： key1,key2,key3
func (this *AwsCfg)DynamoDbGetItemByKvs(table string, kvs map[string]string, ProjectionExpression string)(*dynamodb.GetItemOutput, error){
	tableName := this.DynamoDbCfg.TablePrefix + table
	params := &dynamodb.GetItemInput{
		TableName: &tableName,
	}
	if ProjectionExpression != "" {
		params.ProjectionExpression = aws.String(ProjectionExpression)
	}

	aKey := map[string]*dynamodb.AttributeValue{}
	for key, val := range kvs {
		tmp := dynamodb.AttributeValue{}
		tmp.S = aws.String(val)
		aKey[key] = &tmp
	}
	params.Key = aKey
	result, err := this.DynamoDbClient.GetItem(params)
	return result, err
}


//测试更新dynamodb
func (this *AwsCfg)DynamoDbUpdateItemTest()(*dynamodb.UpdateItemOutput, error){
	type SaveInfo struct {
		SaveId	string	`json:"saveId"`
	}
	type updateSaveInfo struct {
		Ver		int		`json:":ver"`
		Fver	int		`json:":fver"`
	}
	aKey := SaveInfo{
		SaveId: "1209908",
	}
	updateSave := updateSaveInfo{
		Fver: 410,
		Ver: 1606967077,
	}
	ret, err := this.DynamoDbUpdateItem("SaveInfo", aKey, updateSave)
	fmt.Println("DynamoDbUpdateItem : ", ret, err)
	return ret, err
}


//更新一条数据
//@param table 表名
//@param aKey 需要更新的条件 结构体形式 主键key Map【非二级索引】
//@param aUpdate 需要更新的值 结构体形式 json tag 需要有":"相关的更新字段
//@param UpdateExpression 需要更新的字段，结构体中的json字段如："set ver = :ver"
func (this *AwsCfg)DynamoDbUpdateItem(table string, aKey interface{}, aUpdate interface{})(*dynamodb.UpdateItemOutput, error){
	keyMap, err := dynamodbattribute.MarshalMap(aKey)
	if err != nil {
		return nil, err
	}
	updateMap, err := dynamodbattribute.MarshalMap(aUpdate)
	if err != nil {
		return nil, err
	}

	//根据反射获取结构的数据
	updateType := reflect.TypeOf(aUpdate)
	num := updateType.NumField()
	var updateField []string
	for i:=0; i < num; i++{
		jsonStr := updateType.Field(i).Tag.Get("json")
		field := strings.Trim(jsonStr, ":")
		updateField = append(updateField, field + "=" + jsonStr)
	}
	updateString := "set " + strings.Join(updateField, ",")
	tableName := this.DynamoDbCfg.TablePrefix + table
	params := &dynamodb.UpdateItemInput{
		TableName: aws.String(tableName),
		ReturnValues: aws.String("UPDATED_NEW"),
		Key: keyMap,
		ExpressionAttributeValues: updateMap,
		UpdateExpression: aws.String(updateString),//如： set ver=:ver,fver=:fver
	}
	result, err := this.DynamoDbClient.UpdateItem(params)
	return result, err
}