package common // import "github.com/starquantum/gocy/common"


FUNCTIONS

func CopyFile(dstName, srcName string) error
    复制文件

func CutFile(dstName, srcName string) error
    剪切文件

func FFExec(args ...string) (gjson.Result, error)
    调用ffmpeg执行命令并返回结果 依赖ffmpeg

func FormatSubtitleTime(duration float64) string
    将时长格式化为字幕时长格式 39.77 >> 00:00:39,770

func FormatTimeStr(t time.Time) string
    格式化事件为标准的字符串

func GetFileParam(path string, codec_type string) gjson.Result
    解析音视频文件参数 依赖 ffprobe

func IsNullOrEmpty(str string) bool
    是否空字符

func Isfile(filename string) bool
    文件是否存在

func MakePath(filePath string) error
    创建文件夹，文件夹如果存在 直接返回

func Md5BufHex(buf []byte) string
    获取数据md5的hex编码

func Md5FileHex(path string) string
    获取文件md5的hex编码

func PathExists(path string) (bool, error)
    文件文件或文件夹是否存在

func PureFilename(name string) string
    获取纯净文件名 用于保存到磁盘本地

func RefreshAccessTime(filePath string) error
    更新文件访问时间为当前时间

func TrimAllSpace(str string) string
    删除字符串中的所有空格符

func TrimEnter(str string) string
    删除换行符

func Unzip(zipFile string, destDir string) error
    解压zip文件

func UuidV1() string
    uuidv1

func UuidV4() string
    uuidv4

func ZipArchive(source, target, filter string) error
    压缩为zip格式 source为要压缩的文件或文件夹, 绝对路径和相对路径都可以 target是目标文件 filter是过滤正则(Golang 的 包
    path.Match)

func ZipArchiveFiles(files map[string]string, target string) error
    打包指定的文件集合

func NewFileLock(dir string) *fileLock

TYPES

type FileStat struct {
	FileInfo os.FileInfo
	// Has unexported fields.
}
    文件状态信息

func Stat(filePath string) (*FileStat, error)
    查询文件状态信息

func (fs *FileStat) ChangeAccessTime(timestamp int64) error
    修改文件访问时间

func (fs *FileStat) GetAccessTime() int64
    获取文件的访问时间, 返回时间戳

func (fs *FileStat) GetCreateTime() int64
    获取文件的创建时间, 返回时间戳

func (fs *FileStat) GetWriteTime() int64
    获取文件的修改时间, 返回时间戳

